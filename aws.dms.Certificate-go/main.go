package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewCertificate(ctx, "test", &dms.CertificateArgs{
			CertificateId:  pulumi.String("test-dms-certificate-tf"),
			CertificatePem: pulumi.String("..."),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

