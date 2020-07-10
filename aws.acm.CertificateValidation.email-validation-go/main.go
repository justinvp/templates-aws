package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		certCertificate, err := acm.NewCertificate(ctx, "certCertificate", &acm.CertificateArgs{
			DomainName:       pulumi.String("example.com"),
			ValidationMethod: pulumi.String("EMAIL"),
		})
		if err != nil {
			return err
		}
		_, err = acm.NewCertificateValidation(ctx, "certCertificateValidation", &acm.CertificateValidationArgs{
			CertificateArn: certCertificate.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

