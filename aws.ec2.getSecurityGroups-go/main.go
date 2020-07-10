package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.GetSecurityGroups(ctx, &ec2.GetSecurityGroupsArgs{
			Tags: map[string]interface{}{
				"Application": "k8s",
				"Environment": "dev",
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

