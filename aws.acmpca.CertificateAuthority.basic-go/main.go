package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acmpca"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := acmpca.NewCertificateAuthority(ctx, "example", &acmpca.CertificateAuthorityArgs{
			CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
				KeyAlgorithm:     pulumi.String("RSA_4096"),
				SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
				Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
					CommonName: pulumi.String("example.com"),
				},
			},
			PermanentDeletionTimeInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

