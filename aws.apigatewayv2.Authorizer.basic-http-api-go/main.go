package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewAuthorizer(ctx, "example", &apigatewayv2.AuthorizerArgs{
			ApiId:          pulumi.Any(aws_apigatewayv2_api.Example.Id),
			AuthorizerType: pulumi.String("JWT"),
			IdentitySources: pulumi.StringArray{
				pulumi.String(fmt.Sprintf("%v%v", "$", "request.header.Authorization")),
			},
			JwtConfiguration: &apigatewayv2.AuthorizerJwtConfigurationArgs{
				Audience: pulumi.StringArray{
					pulumi.String("example"),
				},
				Issuer: pulumi.String(fmt.Sprintf("%v%v", "https://", aws_cognito_user_pool.Example.Endpoint)),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

