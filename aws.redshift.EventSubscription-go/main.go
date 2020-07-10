package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultCluster, err := redshift.NewCluster(ctx, "defaultCluster", &redshift.ClusterArgs{
			ClusterIdentifier: pulumi.String("default"),
			DatabaseName:      pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
		if err != nil {
			return err
		}
		_, err = redshift.NewEventSubscription(ctx, "defaultEventSubscription", &redshift.EventSubscriptionArgs{
			EventCategories: pulumi.StringArray{
				pulumi.String("configuration"),
				pulumi.String("management"),
				pulumi.String("monitoring"),
				pulumi.String("security"),
			},
			Severity:    pulumi.String("INFO"),
			SnsTopicArn: defaultTopic.Arn,
			SourceIds: pulumi.StringArray{
				defaultCluster.ID(),
			},
			SourceType: pulumi.String("cluster"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("default"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

