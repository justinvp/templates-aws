package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewAuthorizer(ctx, "example", &apigatewayv2.AuthorizerArgs{
			ApiId:          pulumi.Any(aws_apigatewayv2_api.Example.Id),
			AuthorizerType: pulumi.String("REQUEST"),
			AuthorizerUri:  pulumi.Any(aws_lambda_function.Example.Invoke_arn),
			IdentitySources: pulumi.StringArray{
				pulumi.String("route.request.header.Auth"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

