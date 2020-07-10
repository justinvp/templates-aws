package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acm"
	"github.com/pulumi/pulumi-tls/sdk/v2/go/tls"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplePrivateKey, err := tls.NewPrivateKey(ctx, "examplePrivateKey", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
		})
		if err != nil {
			return err
		}
		exampleSelfSignedCert, err := tls.NewSelfSignedCert(ctx, "exampleSelfSignedCert", &tls.SelfSignedCertArgs{
			AllowedUses: pulumi.StringArray{
				pulumi.String("key_encipherment"),
				pulumi.String("digital_signature"),
				pulumi.String("server_auth"),
			},
			KeyAlgorithm:  pulumi.String("RSA"),
			PrivateKeyPem: examplePrivateKey.PrivateKeyPem,
			Subjects: tls.SelfSignedCertSubjectArray{
				&tls.SelfSignedCertSubjectArgs{
					CommonName:   pulumi.String("example.com"),
					Organization: pulumi.String("ACME Examples, Inc"),
				},
			},
			ValidityPeriodHours: pulumi.Int(12),
		})
		if err != nil {
			return err
		}
		_, err = acm.NewCertificate(ctx, "cert", &acm.CertificateArgs{
			CertificateBody: exampleSelfSignedCert.CertPem,
			PrivateKey:      examplePrivateKey.PrivateKeyPem,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

