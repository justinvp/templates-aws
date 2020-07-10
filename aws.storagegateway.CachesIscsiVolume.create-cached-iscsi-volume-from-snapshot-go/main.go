package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.NewCachesIscsiVolume(ctx, "example", &storagegateway.CachesIscsiVolumeArgs{
			GatewayArn:         pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
			NetworkInterfaceId: pulumi.Any(aws_instance.Example.Private_ip),
			SnapshotId:         pulumi.Any(aws_ebs_snapshot.Example.Id),
			TargetName:         pulumi.String("example"),
			VolumeSizeInBytes:  pulumi.Int(aws_ebs_snapshot.Example.Volume_size * 1024 * 1024 * 1024),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

