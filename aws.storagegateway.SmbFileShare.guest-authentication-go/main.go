package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.NewSmbFileShare(ctx, "example", &storagegateway.SmbFileShareArgs{
			Authentication: pulumi.String("GuestAccess"),
			GatewayArn:     pulumi.Any(aws_storagegateway_gateway.Example.Arn),
			LocationArn:    pulumi.Any(aws_s3_bucket.Example.Arn),
			RoleArn:        pulumi.Any(aws_iam_role.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

