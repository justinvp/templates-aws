package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewLogSubscriptionFilter(ctx, "testLambdafunctionLogfilter", &cloudwatch.LogSubscriptionFilterArgs{
			DestinationArn: pulumi.Any(aws_kinesis_stream.Test_logstream.Arn),
			Distribution:   pulumi.String("Random"),
			FilterPattern:  pulumi.String("logtype test"),
			LogGroup:       pulumi.String("/aws/lambda/example_lambda_name"),
			RoleArn:        pulumi.Any(aws_iam_role.Iam_for_lambda.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

