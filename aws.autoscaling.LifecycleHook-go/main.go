package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foobarGroup, err := autoscaling.NewGroup(ctx, "foobarGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
			},
			HealthCheckType: pulumi.String("EC2"),
			Tags: autoscaling.GroupTagArray{
				&autoscaling.GroupTagArgs{
					Key:               pulumi.String("Foo"),
					PropagateAtLaunch: pulumi.Bool(true),
					Value:             pulumi.String("foo-bar"),
				},
			},
			TerminationPolicies: pulumi.StringArray{
				pulumi.String("OldestInstance"),
			},
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewLifecycleHook(ctx, "foobarLifecycleHook", &autoscaling.LifecycleHookArgs{
			AutoscalingGroupName:  foobarGroup.Name,
			DefaultResult:         pulumi.String("CONTINUE"),
			HeartbeatTimeout:      pulumi.Int(2000),
			LifecycleTransition:   pulumi.String("autoscaling:EC2_INSTANCE_LAUNCHING"),
			NotificationMetadata:  pulumi.String(fmt.Sprintf("%v%v%v%v", "{\n", "  \"foo\": \"bar\"\n", "}\n", "\n")),
			NotificationTargetArn: pulumi.String("arn:aws:sqs:us-east-1:444455556666:queue1*"),
			RoleArn:               pulumi.String("arn:aws:iam::123456789012:role/S3Access"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

