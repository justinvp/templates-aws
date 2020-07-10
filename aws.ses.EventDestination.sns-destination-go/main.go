package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewEventDestination(ctx, "sns", &ses.EventDestinationArgs{
			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
			Enabled:              pulumi.Bool(true),
			MatchingTypes: pulumi.StringArray{
				pulumi.String("bounce"),
				pulumi.String("send"),
			},
			SnsDestination: &ses.EventDestinationSnsDestinationArgs{
				TopicArn: pulumi.Any(aws_sns_topic.Example.Arn),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

