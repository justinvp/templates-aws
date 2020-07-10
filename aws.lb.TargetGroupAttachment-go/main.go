package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testTargetGroup, err := lb.NewTargetGroup(ctx, "testTargetGroup", nil)
		if err != nil {
			return err
		}
		testInstance, err := ec2.NewInstance(ctx, "testInstance", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewTargetGroupAttachment(ctx, "testTargetGroupAttachment", &lb.TargetGroupAttachmentArgs{
			Port:           pulumi.Int(80),
			TargetGroupArn: testTargetGroup.Arn,
			TargetId:       testInstance.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

