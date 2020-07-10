package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := aws_volume_attachment.Test.Device_name
		_, err := storagegateway.GetLocalDisk(ctx, &storagegateway.GetLocalDiskArgs{
			DiskPath:   &opt0,
			GatewayArn: aws_storagegateway_gateway.Test.Arn,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

