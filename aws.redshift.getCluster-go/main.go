package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCluster, err := redshift.LookupCluster(ctx, &redshift.LookupClusterArgs{
			ClusterIdentifier: "test-cluster",
		}, nil)
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("redshift"),
			RedshiftConfiguration: &kinesis.FirehoseDeliveryStreamRedshiftConfigurationArgs{
				ClusterJdbcurl:   pulumi.String(fmt.Sprintf("%v%v%v%v", "jdbc:redshift://", testCluster.Endpoint, "/", testCluster.DatabaseName)),
				CopyOptions:      pulumi.String("delimiter '|'"),
				DataTableColumns: pulumi.String("test-col"),
				DataTableName:    pulumi.String("test-table"),
				Password:         pulumi.String("T3stPass"),
				RoleArn:          pulumi.Any(aws_iam_role.Firehose_role.Arn),
				Username:         pulumi.String("testuser"),
			},
			S3Configuration: &kinesis.FirehoseDeliveryStreamS3ConfigurationArgs{
				BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
				BufferInterval:    pulumi.Int(400),
				BufferSize:        pulumi.Int(10),
				CompressionFormat: pulumi.String("GZIP"),
				RoleArn:           pulumi.Any(aws_iam_role.Firehose_role.Arn),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

