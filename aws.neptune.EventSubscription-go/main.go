package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/neptune"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultCluster, err := neptune.NewCluster(ctx, "defaultCluster", &neptune.ClusterArgs{
			ApplyImmediately:                 pulumi.Bool(true),
			BackupRetentionPeriod:            pulumi.Int(5),
			ClusterIdentifier:                pulumi.String("neptune-cluster-demo"),
			Engine:                           pulumi.String("neptune"),
			IamDatabaseAuthenticationEnabled: pulumi.Bool(true),
			PreferredBackupWindow:            pulumi.String("07:00-09:00"),
			SkipFinalSnapshot:                pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		example, err := neptune.NewClusterInstance(ctx, "example", &neptune.ClusterInstanceArgs{
			ApplyImmediately:  pulumi.Bool(true),
			ClusterIdentifier: defaultCluster.ID(),
			Engine:            pulumi.String("neptune"),
			InstanceClass:     pulumi.String("db.r4.large"),
		})
		if err != nil {
			return err
		}
		defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
		if err != nil {
			return err
		}
		_, err = neptune.NewEventSubscription(ctx, "defaultEventSubscription", &neptune.EventSubscriptionArgs{
			EventCategories: pulumi.StringArray{
				pulumi.String("maintenance"),
				pulumi.String("availability"),
				pulumi.String("creation"),
				pulumi.String("backup"),
				pulumi.String("restoration"),
				pulumi.String("recovery"),
				pulumi.String("deletion"),
				pulumi.String("failover"),
				pulumi.String("failure"),
				pulumi.String("notification"),
				pulumi.String("configuration change"),
				pulumi.String("read replica"),
			},
			SnsTopicArn: defaultTopic.Arn,
			SourceIds: pulumi.StringArray{
				example.ID(),
			},
			SourceType: pulumi.String("db-instance"),
			Tags: pulumi.StringMap{
				"env": pulumi.String("test"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

