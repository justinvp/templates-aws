package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewAmiLaunchPermission(ctx, "example", &ec2.AmiLaunchPermissionArgs{
			AccountId: pulumi.String("123456789012"),
			ImageId:   pulumi.String("ami-12345678"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

