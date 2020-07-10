package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewAnalyticsApplication(ctx, "testApplication", &kinesis.AnalyticsApplicationArgs{
			Inputs: &kinesis.AnalyticsApplicationInputsArgs{
				KinesisStream: &kinesis.AnalyticsApplicationInputsKinesisStreamArgs{
					ResourceArn: testStream.Arn,
					RoleArn:     pulumi.Any(aws_iam_role.Test.Arn),
				},
				NamePrefix: pulumi.String("test_prefix"),
				Parallelism: &kinesis.AnalyticsApplicationInputsParallelismArgs{
					Count: pulumi.Int(1),
				},
				Schema: &kinesis.AnalyticsApplicationInputsSchemaArgs{
					RecordColumns: kinesis.AnalyticsApplicationInputsSchemaRecordColumnArray{
						&kinesis.AnalyticsApplicationInputsSchemaRecordColumnArgs{
							Mapping: pulumi.String(fmt.Sprintf("%v%v", "$", ".test")),
							Name:    pulumi.String("test"),
							SqlType: pulumi.String("VARCHAR(8)"),
						},
					},
					RecordEncoding: pulumi.String("UTF-8"),
					RecordFormat: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatArgs{
						MappingParameters: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs{
							Json: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs{
								RecordRowPath: pulumi.String("$"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

