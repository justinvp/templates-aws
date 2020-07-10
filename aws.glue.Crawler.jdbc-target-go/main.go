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
			JdbcTargets: glue.CrawlerJdbcTargetArray{
				&glue.CrawlerJdbcTargetArgs{
					ConnectionName: pulumi.Any(aws_glue_connection.Example.Name),
					Path:           pulumi.String(fmt.Sprintf("%v%v", "database-name/", "%")),
				},
			},
			Role: pulumi.Any(aws_iam_role.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

