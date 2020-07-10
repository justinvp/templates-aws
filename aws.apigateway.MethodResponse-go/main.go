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
			ParentId: myDemoAPI.RootResourceId,
			PathPart: pulumi.String("mydemoresource"),
			RestApi:  myDemoAPI.ID(),
		})
		if err != nil {
			return err
		}
		myDemoMethod, err := apigateway.NewMethod(ctx, "myDemoMethod", &apigateway.MethodArgs{
			Authorization: pulumi.String("NONE"),
			HttpMethod:    pulumi.String("GET"),
			ResourceId:    myDemoResource.ID(),
			RestApi:       myDemoAPI.ID(),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewIntegration(ctx, "myDemoIntegration", &apigateway.IntegrationArgs{
			HttpMethod: myDemoMethod.HttpMethod,
			ResourceId: myDemoResource.ID(),
			RestApi:    myDemoAPI.ID(),
			Type:       pulumi.String("MOCK"),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewMethodResponse(ctx, "response200", &apigateway.MethodResponseArgs{
			HttpMethod: myDemoMethod.HttpMethod,
			ResourceId: myDemoResource.ID(),
			RestApi:    myDemoAPI.ID(),
			StatusCode: pulumi.String("200"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

