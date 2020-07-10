package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewDeployment(ctx, "example", &apigatewayv2.DeploymentArgs{
			ApiId:       pulumi.Any(aws_apigatewayv2_route.Example.Api_id),
			Description: pulumi.String("Example deployment"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

