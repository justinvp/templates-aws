package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := sns.NewTopic(ctx, "example", nil)
		if err != nil {
			return err
		}
		bar, err := autoscaling.NewGroup(ctx, "bar", nil)
		if err != nil {
			return err
		}
		foo, err := autoscaling.NewGroup(ctx, "foo", nil)
		if err != nil {
			return err
		}
		_, err = autoscaling.NewNotification(ctx, "exampleNotifications", &autoscaling.NotificationArgs{
			GroupNames: pulumi.StringArray{
				bar.Name,
				foo.Name,
			},
			Notifications: pulumi.StringArray{
				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH"),
				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE"),
				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH_ERROR"),
				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE_ERROR"),
			},
			TopicArn: example.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

