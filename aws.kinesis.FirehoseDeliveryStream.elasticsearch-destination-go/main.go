package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticsearch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCluster, err := elasticsearch.NewDomain(ctx, "testCluster", nil)
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("elasticsearch"),
			ElasticsearchConfiguration: &kinesis.FirehoseDeliveryStreamElasticsearchConfigurationArgs{
				DomainArn: testCluster.Arn,
				IndexName: pulumi.String("test"),
				ProcessingConfiguration: &kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationArgs{
					Enabled: pulumi.Bool(true),
					Processors: kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorArray{
						&kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorArgs{
							Parameters: kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorParameterArray{
								&kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("LambdaArn"),
									ParameterValue: pulumi.String(fmt.Sprintf("%v%v%v%v", aws_lambda_function.Lambda_processor.Arn, ":", "$", "LATEST")),
								},
							},
							Type: pulumi.String("Lambda"),
						},
					},
				},
				RoleArn:  pulumi.Any(aws_iam_role.Firehose_role.Arn),
				TypeName: pulumi.String("test"),
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

