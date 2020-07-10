package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		groups, err := aws.GetAutoscalingGroups(ctx, &aws.GetAutoscalingGroupsArgs{
			Filters: []aws.GetAutoscalingGroupsFilter{
				aws.GetAutoscalingGroupsFilter{
					Name: "key",
					Values: []string{
						"Team",
					},
				},
				aws.GetAutoscalingGroupsFilter{
					Name: "value",
					Values: []string{
						"Pets",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = autoscaling.NewNotification(ctx, "slackNotifications", &autoscaling.NotificationArgs{
			GroupNames: toPulumiStringArray(groups.Names),
			Notifications: pulumi.StringArray{
				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH"),
				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE"),
				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH_ERROR"),
				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE_ERROR"),
			},
			TopicArn: pulumi.String("TOPIC ARN"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
func toPulumiStringArray(arr []string) pulumi.StringArray {
	var pulumiArr pulumi.StringArray
	for _, v := range arr {
		pulumiArr = append(pulumiArr, pulumi.String(v))
	}
	return pulumiArr
}

