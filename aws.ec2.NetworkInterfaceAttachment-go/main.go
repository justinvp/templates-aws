package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewNetworkInterfaceAttachment(ctx, "test", &ec2.NetworkInterfaceAttachmentArgs{
			DeviceIndex:        pulumi.Int(0),
			InstanceId:         pulumi.Any(aws_instance.Test.Id),
			NetworkInterfaceId: pulumi.Any(aws_network_interface.Test.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

