package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewApiMapping(ctx, "example", &apigatewayv2.ApiMappingArgs{
			ApiId:      pulumi.Any(aws_apigatewayv2_api.Example.Id),
			DomainName: pulumi.Any(aws_apigatewayv2_domain_name.Example.Id),
			Stage:      pulumi.Any(aws_apigatewayv2_stage.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

