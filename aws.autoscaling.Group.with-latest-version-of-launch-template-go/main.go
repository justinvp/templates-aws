package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foobar, err := ec2.NewLaunchTemplate(ctx, "foobar", &ec2.LaunchTemplateArgs{
			ImageId:      pulumi.String("ami-1a2b3c"),
			InstanceType: pulumi.String("t2.micro"),
			NamePrefix:   pulumi.String("foobar"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			DesiredCapacity: pulumi.Int(1),
			LaunchTemplate: &autoscaling.GroupLaunchTemplateArgs{
				Id:      foobar.ID(),
				Version: pulumi.String(fmt.Sprintf("%v%v", "$", "Latest")),
			},
			MaxSize: pulumi.Int(1),
			MinSize: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

