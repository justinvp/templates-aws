package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cur"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cur.NewReportDefinition(ctx, "exampleCurReportDefinition", &cur.ReportDefinitionArgs{
			AdditionalArtifacts: pulumi.StringArray{
				pulumi.String("REDSHIFT"),
				pulumi.String("QUICKSIGHT"),
			},
			AdditionalSchemaElements: pulumi.StringArray{
				pulumi.String("RESOURCES"),
			},
			Compression: pulumi.String("GZIP"),
			Format:      pulumi.String("textORcsv"),
			ReportName:  pulumi.String("example-cur-report-definition"),
			S3Bucket:    pulumi.String("example-bucket-name"),
			S3Region:    pulumi.String("us-east-1"),
			TimeUnit:    pulumi.String("HOURLY"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

