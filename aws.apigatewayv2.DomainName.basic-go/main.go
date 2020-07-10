package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewDomainName(ctx, "example", &apigatewayv2.DomainNameArgs{
			DomainName: pulumi.String("ws-api.example.com"),
			DomainNameConfiguration: &apigatewayv2.DomainNameDomainNameConfigurationArgs{
				CertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
				EndpointType:   pulumi.String("REGIONAL"),
				SecurityPolicy: pulumi.String("TLS_1_2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

