package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewTrafficMirrorTarget(ctx, "nlb", &ec2.TrafficMirrorTargetArgs{
			Description:            pulumi.String("NLB target"),
			NetworkLoadBalancerArn: pulumi.Any(aws_lb.Lb.Arn),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewTrafficMirrorTarget(ctx, "eni", &ec2.TrafficMirrorTargetArgs{
			Description:        pulumi.String("ENI target"),
			NetworkInterfaceId: pulumi.Any(aws_instance.Test.Primary_network_interface_id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

