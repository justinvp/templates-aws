package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewEventDestination(ctx, "kinesis", &ses.EventDestinationArgs{
			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
			Enabled:              pulumi.Bool(true),
			KinesisDestination: &ses.EventDestinationKinesisDestinationArgs{
				RoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
				StreamArn: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
			},
			MatchingTypes: pulumi.StringArray{
				pulumi.String("bounce"),
				pulumi.String("send"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

