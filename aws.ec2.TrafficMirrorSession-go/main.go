package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		filter, err := ec2.NewTrafficMirrorFilter(ctx, "filter", &ec2.TrafficMirrorFilterArgs{
			Description: pulumi.String("traffic mirror filter - example"),
			NetworkServices: pulumi.StringArray{
				pulumi.String("amazon-dns"),
			},
		})
		if err != nil {
			return err
		}
		target, err := ec2.NewTrafficMirrorTarget(ctx, "target", &ec2.TrafficMirrorTargetArgs{
			NetworkLoadBalancerArn: pulumi.Any(aws_lb.Lb.Arn),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewTrafficMirrorSession(ctx, "session", &ec2.TrafficMirrorSessionArgs{
			Description:           pulumi.String("traffic mirror session - example"),
			NetworkInterfaceId:    pulumi.Any(aws_instance.Test.Primary_network_interface_id),
			TrafficMirrorFilterId: filter.ID(),
			TrafficMirrorTargetId: target.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

