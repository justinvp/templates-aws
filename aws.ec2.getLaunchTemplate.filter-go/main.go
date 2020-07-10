package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.LookupLaunchTemplate(ctx, &ec2.LookupLaunchTemplateArgs{
			Filters: []ec2.GetLaunchTemplateFilter{
				ec2.GetLaunchTemplateFilter{
					Name: "launch-template-name",
					Values: []string{
						"some-template",
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

