package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewRoute(ctx, "example", &apigatewayv2.RouteArgs{
			ApiId:    pulumi.Any(aws_apigatewayv2_api.Example.Id),
			RouteKey: pulumi.String(fmt.Sprintf("%v%v", "$", "default")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

