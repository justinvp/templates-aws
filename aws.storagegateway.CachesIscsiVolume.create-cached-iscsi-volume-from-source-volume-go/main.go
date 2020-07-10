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
			SourceVolumeArn:    pulumi.Any(aws_storagegateway_cached_iscsi_volume.Existing.Arn),
			TargetName:         pulumi.String("example"),
			VolumeSizeInBytes:  pulumi.Any(aws_storagegateway_cached_iscsi_volume.Existing.Volume_size_in_bytes),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

