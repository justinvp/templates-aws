package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultInstance, err := rds.NewInstance(ctx, "defaultInstance", &rds.InstanceArgs{
			AllocatedStorage:   pulumi.Int(10),
			DbSubnetGroupName:  pulumi.String("my_database_subnet_group"),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("5.6.17"),
			InstanceClass:      pulumi.String("db.t2.micro"),
			Name:               pulumi.String("mydb"),
			ParameterGroupName: pulumi.String("default.mysql5.6"),
			Password:           pulumi.String("bar"),
			Username:           pulumi.String("foo"),
		})
		if err != nil {
			return err
		}
		defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
		if err != nil {
			return err
		}
		_, err = rds.NewEventSubscription(ctx, "defaultEventSubscription", &rds.EventSubscriptionArgs{
			EventCategories: pulumi.StringArray{
				pulumi.String("availability"),
				pulumi.String("deletion"),
				pulumi.String("failover"),
				pulumi.String("failure"),
				pulumi.String("low storage"),
				pulumi.String("maintenance"),
				pulumi.String("notification"),
				pulumi.String("read replica"),
				pulumi.String("recovery"),
				pulumi.String("restoration"),
			},
			SnsTopic: defaultTopic.Arn,
			SourceIds: pulumi.StringArray{
				defaultInstance.ID(),
			},
			SourceType: pulumi.String("db-instance"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

