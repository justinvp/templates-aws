package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autoscaling.NewAttachment(ctx, "asgAttachmentBar", &autoscaling.AttachmentArgs{
			AutoscalingGroupName: pulumi.Any(aws_autoscaling_group.Asg.Id),
			Elb:                  pulumi.Any(aws_elb.Bar.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

