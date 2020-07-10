package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetAmiIds(ctx, &aws.GetAmiIdsArgs{
			Filters: []aws.GetAmiIdsFilter{
				aws.GetAmiIdsFilter{
					Name: "name",
					Values: []string{
						"ubuntu/images/ubuntu-*-*-amd64-server-*",
					},
				},
			},
			Owners: []string{
				"099720109477",
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

