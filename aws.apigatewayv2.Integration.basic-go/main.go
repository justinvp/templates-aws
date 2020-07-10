package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewIntegration(ctx, "example", &apigatewayv2.IntegrationArgs{
			ApiId:           pulumi.Any(aws_apigatewayv2_api.Example.Id),
			IntegrationType: pulumi.String("MOCK"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

