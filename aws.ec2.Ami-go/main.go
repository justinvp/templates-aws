package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewAmi(ctx, "example", &ec2.AmiArgs{
			EbsBlockDevices: ec2.AmiEbsBlockDeviceArray{
				&ec2.AmiEbsBlockDeviceArgs{
					DeviceName: pulumi.String("/dev/xvda"),
					SnapshotId: pulumi.String("snap-xxxxxxxx"),
					VolumeSize: pulumi.Int(8),
				},
			},
			RootDeviceName:     pulumi.String("/dev/xvda"),
			VirtualizationType: pulumi.String("hvm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

