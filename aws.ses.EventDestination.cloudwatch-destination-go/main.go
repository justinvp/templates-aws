package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewEventDestination(ctx, "cloudwatch", &ses.EventDestinationArgs{
			CloudwatchDestinations: ses.EventDestinationCloudwatchDestinationArray{
				&ses.EventDestinationCloudwatchDestinationArgs{
					DefaultValue:  pulumi.String("default"),
					DimensionName: pulumi.String("dimension"),
					ValueSource:   pulumi.String("emailHeader"),
				},
			},
			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
			Enabled:              pulumi.Bool(true),
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

