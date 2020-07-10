package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.NewAccountPublicAccessBlock(ctx, "example", &s3.AccountPublicAccessBlockArgs{
			BlockPublicAcls:   pulumi.Bool(true),
			BlockPublicPolicy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

