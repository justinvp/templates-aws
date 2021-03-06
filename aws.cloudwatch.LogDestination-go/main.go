package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewLogDestination(ctx, "testDestination", &cloudwatch.LogDestinationArgs{
			RoleArn:   pulumi.Any(aws_iam_role.Iam_for_cloudwatch.Arn),
			TargetArn: pulumi.Any(aws_kinesis_stream.Kinesis_for_cloudwatch.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

