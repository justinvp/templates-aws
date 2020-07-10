package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		console, err := cloudwatch.NewEventRule(ctx, "console", &cloudwatch.EventRuleArgs{
			Description:  pulumi.String("Capture all EC2 scaling events"),
			EventPattern: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"source\": [\n", "    \"aws.autoscaling\"\n", "  ],\n", "  \"detail-type\": [\n", "    \"EC2 Instance Launch Successful\",\n", "    \"EC2 Instance Terminate Successful\",\n", "    \"EC2 Instance Launch Unsuccessful\",\n", "    \"EC2 Instance Terminate Unsuccessful\"\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "yada", &cloudwatch.EventTargetArgs{
			Arn:  testStream.Arn,
			Rule: console.Name,
			RunCommandTargets: cloudwatch.EventTargetRunCommandTargetArray{
				&cloudwatch.EventTargetRunCommandTargetArgs{
					Key: pulumi.String("tag:Name"),
					Values: pulumi.StringArray{
						pulumi.String("FooBar"),
					},
				},
				&cloudwatch.EventTargetRunCommandTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						pulumi.String("i-162058cd308bffec2"),
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

