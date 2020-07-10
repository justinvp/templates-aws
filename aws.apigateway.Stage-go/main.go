package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testRestApi, err := apigateway.NewRestApi(ctx, "testRestApi", &apigateway.RestApiArgs{
			Description: pulumi.String("This is my API for demonstration purposes"),
		})
		if err != nil {
			return err
		}
		testDeployment, err := apigateway.NewDeployment(ctx, "testDeployment", &apigateway.DeploymentArgs{
			RestApi:   testRestApi.ID(),
			StageName: pulumi.String("dev"),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_api_gateway_integration.test",
		}))
		if err != nil {
			return err
		}
		testStage, err := apigateway.NewStage(ctx, "testStage", &apigateway.StageArgs{
			Deployment: testDeployment.ID(),
			RestApi:    testRestApi.ID(),
			StageName:  pulumi.String("prod"),
		})
		if err != nil {
			return err
		}
		testResource, err := apigateway.NewResource(ctx, "testResource", &apigateway.ResourceArgs{
			ParentId: testRestApi.RootResourceId,
			PathPart: pulumi.String("mytestresource"),
			RestApi:  testRestApi.ID(),
		})
		if err != nil {
			return err
		}
		testMethod, err := apigateway.NewMethod(ctx, "testMethod", &apigateway.MethodArgs{
			Authorization: pulumi.String("NONE"),
			HttpMethod:    pulumi.String("GET"),
			ResourceId:    testResource.ID(),
			RestApi:       testRestApi.ID(),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewMethodSettings(ctx, "methodSettings", &apigateway.MethodSettingsArgs{
			MethodPath: pulumi.All(testResource.PathPart, testMethod.HttpMethod).ApplyT(func(_args []interface{}) (string, error) {
				pathPart := _args[0].(string)
				httpMethod := _args[1].(string)
				return fmt.Sprintf("%v%v%v", pathPart, "/", httpMethod), nil
			}).(pulumi.StringOutput),
			RestApi: testRestApi.ID(),
			Settings: &apigateway.MethodSettingsSettingsArgs{
				LoggingLevel:   pulumi.String("INFO"),
				MetricsEnabled: pulumi.Bool(true),
			},
			StageName: testStage.StageName,
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewIntegration(ctx, "testIntegration", &apigateway.IntegrationArgs{
			HttpMethod: testMethod.HttpMethod,
			ResourceId: testResource.ID(),
			RestApi:    testRestApi.ID(),
			Type:       pulumi.String("MOCK"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

