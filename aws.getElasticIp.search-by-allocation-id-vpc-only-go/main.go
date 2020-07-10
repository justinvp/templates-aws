package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "eipalloc-12345678"
		_, err := aws.GetElasticIp(ctx, &aws.GetElasticIpArgs{
			Id: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

