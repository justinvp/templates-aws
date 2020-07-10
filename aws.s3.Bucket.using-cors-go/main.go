package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Acl: pulumi.String("public-read"),
			CorsRules: s3.BucketCorsRuleArray{
				&s3.BucketCorsRuleArgs{
					AllowedHeaders: pulumi.StringArray{
						pulumi.String("*"),
					},
					AllowedMethods: pulumi.StringArray{
						pulumi.String("PUT"),
						pulumi.String("POST"),
					},
					AllowedOrigins: pulumi.StringArray{
						pulumi.String("https://s3-website-test.mydomain.com"),
					},
					ExposeHeaders: pulumi.StringArray{
						pulumi.String("ETag"),
					},
					MaxAgeSeconds: pulumi.Int(3000),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

