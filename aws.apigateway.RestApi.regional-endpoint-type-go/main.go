package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.NewRestApi(ctx, "example", &apigateway.RestApiArgs{
			EndpointConfiguration: &apigateway.RestApiEndpointConfigurationArgs{
				Types: pulumi.String("REGIONAL"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

