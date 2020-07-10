package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetElasticIp(ctx, &aws.GetElasticIpArgs{
			Filters: []aws.GetElasticIpFilter{
				aws.GetElasticIpFilter{
					Name: "tag:Name",
					Values: []string{
						"exampleNameTagValue",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

