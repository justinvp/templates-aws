package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "1.2.3.4"
		_, err := aws.GetElasticIp(ctx, &aws.GetElasticIpArgs{
			PublicIp: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

