package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/macie"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := macie.NewS3BucketAssociation(ctx, "example", &macie.S3BucketAssociationArgs{
			BucketName: pulumi.String("tf-macie-example"),
			ClassificationType: &macie.S3BucketAssociationClassificationTypeArgs{
				OneTime: pulumi.String("FULL"),
			},
			Prefix: pulumi.String("data"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

