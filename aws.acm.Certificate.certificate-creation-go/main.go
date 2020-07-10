package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := acm.NewCertificate(ctx, "cert", &acm.CertificateArgs{
			DomainName: pulumi.String("example.com"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
			},
			ValidationMethod: pulumi.String("DNS"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

