package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCluster, err := redshift.NewCluster(ctx, "testCluster", &redshift.ClusterArgs{
			ClusterIdentifier: pulumi.String(fmt.Sprintf("%v%v%v", "tf-redshift-cluster-", "%", "d")),
			ClusterType:       pulumi.String("single-node"),
			DatabaseName:      pulumi.String("test"),
			MasterPassword:    pulumi.String("T3stPass"),
			MasterUsername:    pulumi.String("testuser"),
			NodeType:          pulumi.String("dc1.large"),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("redshift"),
			RedshiftConfiguration: &kinesis.FirehoseDeliveryStreamRedshiftConfigurationArgs{
				ClusterJdbcurl: pulumi.All(testCluster.Endpoint, testCluster.DatabaseName).ApplyT(func(_args []interface{}) (string, error) {
					endpoint := _args[0].(string)
					databaseName := _args[1].(string)
					return fmt.Sprintf("%v%v%v%v", "jdbc:redshift://", endpoint, "/", databaseName), nil
				}).(pulumi.StringOutput),
				CopyOptions:      pulumi.String("delimiter '|'"),
				DataTableColumns: pulumi.String("test-col"),
				DataTableName:    pulumi.String("test-table"),
				Password:         pulumi.String("T3stPass"),
				RoleArn:          pulumi.Any(aws_iam_role.Firehose_role.Arn),
				S3BackupConfiguration: &kinesis.FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationArgs{
					BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
					BufferInterval:    pulumi.Int(300),
					BufferSize:        pulumi.Int(15),
					CompressionFormat: pulumi.String("GZIP"),
					RoleArn:           pulumi.Any(aws_iam_role.Firehose_role.Arn),
				},
				S3BackupMode: pulumi.String("Enabled"),
				Username:     pulumi.String("testuser"),
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

