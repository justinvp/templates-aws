package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.NewRequestValidator(ctx, "example", &apigateway.RequestValidatorArgs{
			RestApi:                   pulumi.Any(aws_api_gateway_rest_api.Example.Id),
			ValidateRequestBody:       pulumi.Bool(true),
			ValidateRequestParameters: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

