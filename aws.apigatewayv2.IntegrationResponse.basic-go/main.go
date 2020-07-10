package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewIntegrationResponse(ctx, "example", &apigatewayv2.IntegrationResponseArgs{
			ApiId:                  pulumi.Any(aws_apigatewayv2_api.Example.Id),
			IntegrationId:          pulumi.Any(aws_apigatewayv2_integration.Example.Id),
			IntegrationResponseKey: pulumi.String("/200/"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

