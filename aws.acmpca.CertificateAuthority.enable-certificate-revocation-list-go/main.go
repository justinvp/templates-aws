package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acmpca"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPolicy(ctx, "exampleBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: exampleBucket.ID(),
			Policy: acmpcaBucketAccess.ApplyT(func(acmpcaBucketAccess iam.GetPolicyDocumentResult) (string, error) {
				return acmpcaBucketAccess.Json, nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = acmpca.NewCertificateAuthority(ctx, "exampleCertificateAuthority", &acmpca.CertificateAuthorityArgs{
			CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
				KeyAlgorithm:     pulumi.String("RSA_4096"),
				SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
				Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
					CommonName: pulumi.String("example.com"),
				},
			},
			RevocationConfiguration: &acmpca.CertificateAuthorityRevocationConfigurationArgs{
				CrlConfiguration: &acmpca.CertificateAuthorityRevocationConfigurationCrlConfigurationArgs{
					CustomCname:      pulumi.String("crl.example.com"),
					Enabled:          pulumi.Bool(true),
					ExpirationInDays: pulumi.Int(7),
					S3BucketName:     exampleBucket.ID(),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_s3_bucket_policy.example",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

