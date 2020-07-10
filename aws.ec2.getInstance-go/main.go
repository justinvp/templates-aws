package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "i-instanceid"
		_, err := ec2.LookupInstance(ctx, &ec2.LookupInstanceArgs{
			Filters: []ec2.GetInstanceFilter{
				ec2.GetInstanceFilter{
					Name: "image-id",
					Values: []string{
						"ami-xxxxxxxx",
					},
				},
				ec2.GetInstanceFilter{
					Name: "tag:Name",
					Values: []string{
						"instance-name-tag",
					},
				},
			},
			InstanceId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

