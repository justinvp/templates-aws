package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ebs.NewEncryptionByDefault(ctx, "example", &ebs.EncryptionByDefaultArgs{
			Enabled: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

