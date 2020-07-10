package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ebs.NewDefaultKmsKey(ctx, "example", &ebs.DefaultKmsKeyArgs{
			KeyArn: pulumi.Any(aws_kms_key.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

