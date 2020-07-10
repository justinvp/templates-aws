package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.NewCache(ctx, "example", &storagegateway.CacheArgs{
			DiskId:     pulumi.Any(data.Aws_storagegateway_local_disk.Example.Id),
			GatewayArn: pulumi.Any(aws_storagegateway_gateway.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

