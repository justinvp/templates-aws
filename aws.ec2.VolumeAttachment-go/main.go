package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		web, err := ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
			Ami:              pulumi.String("ami-21f78e11"),
			AvailabilityZone: pulumi.String("us-west-2a"),
			InstanceType:     pulumi.String("t1.micro"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Size:             pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVolumeAttachment(ctx, "ebsAtt", &ec2.VolumeAttachmentArgs{
			DeviceName: pulumi.String("/dev/sdh"),
			InstanceId: web.ID(),
			VolumeId:   example.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

