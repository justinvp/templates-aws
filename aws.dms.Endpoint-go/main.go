package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewEndpoint(ctx, "test", &dms.EndpointArgs{
			CertificateArn:            pulumi.String("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"),
			DatabaseName:              pulumi.String("test"),
			EndpointId:                pulumi.String("test-dms-endpoint-tf"),
			EndpointType:              pulumi.String("source"),
			EngineName:                pulumi.String("aurora"),
			ExtraConnectionAttributes: pulumi.String(""),
			KmsKeyArn:                 pulumi.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012"),
			Password:                  pulumi.String("test"),
			Port:                      pulumi.Int(3306),
			ServerName:                pulumi.String("test"),
			SslMode:                   pulumi.String("none"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
			},
			Username: pulumi.String("test"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

