package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetElasticIp(ctx, &aws.GetElasticIpArgs{
			Tags: map[string]interface{}{
				"Name": "exampleNameTagValue",
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

