package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
			Description: pulumi.String("This is my API for demonstration purposes"),
		})
		if err != nil {
			return err
		}
		myDemoResource, err := apigateway.NewResource(ctx, "myDemoResource", &apigateway.ResourceArgs{
			RestApi:  myDemoAPI.ID(),
			ParentId: myDemoAPI.RootResourceId,
			PathPart: pulumi.String("test"),
		})
		if err != nil {
			return err
		}
		myDemoMethod, err := apigateway.NewMethod(ctx, "myDemoMethod", &apigateway.MethodArgs{
			RestApi:       myDemoAPI.ID(),
			ResourceId:    myDemoResource.ID(),
			HttpMethod:    pulumi.String("GET"),
			Authorization: pulumi.String("NONE"),
		})
		if err != nil {
			return err
		}
		myDemoIntegration, err := apigateway.NewIntegration(ctx, "myDemoIntegration", &apigateway.IntegrationArgs{
			RestApi:    myDemoAPI.ID(),
			ResourceId: myDemoResource.ID(),
			HttpMethod: myDemoMethod.HttpMethod,
			Type:       pulumi.String("MOCK"),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewDeployment(ctx, "myDemoDeployment", &apigateway.DeploymentArgs{
			RestApi:   myDemoAPI.ID(),
			StageName: pulumi.String("test"),
			Variables: pulumi.StringMap{
				"answer": pulumi.String("42"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			myDemoIntegration,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

