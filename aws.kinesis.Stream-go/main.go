package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
			RetentionPeriod: pulumi.Int(48),
			ShardCount:      pulumi.Int(1),
			ShardLevelMetrics: pulumi.StringArray{
				pulumi.String("IncomingBytes"),
				pulumi.String("OutgoingBytes"),
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

