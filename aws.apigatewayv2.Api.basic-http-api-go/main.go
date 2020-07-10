package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewApi(ctx, "example", &apigatewayv2.ApiArgs{
			ProtocolType: pulumi.String("HTTP"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

