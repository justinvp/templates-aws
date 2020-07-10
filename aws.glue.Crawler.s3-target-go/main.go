package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
			Role:         pulumi.Any(aws_iam_role.Example.Arn),
			S3Targets: glue.CrawlerS3TargetArray{
				&glue.CrawlerS3TargetArgs{
					Path: pulumi.String(fmt.Sprintf("%v%v", "s3://", aws_s3_bucket.Example.Bucket)),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

