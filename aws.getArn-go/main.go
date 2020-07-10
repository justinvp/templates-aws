package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetArn(ctx, &aws.GetArnArgs{
			Arn: "arn:aws:rds:eu-west-1:123456789012:db:mysql-db",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

