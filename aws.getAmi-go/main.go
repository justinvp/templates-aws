package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := true
		opt1 := "^myami-\\d{3}"
		_, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
			ExecutableUsers: []string{
				"self",
			},
			Filters: []aws.GetAmiFilter{
				aws.GetAmiFilter{
					Name: "name",
					Values: []string{
						"myami-*",
					},
				},
				aws.GetAmiFilter{
					Name: "root-device-type",
					Values: []string{
						"ebs",
					},
				},
				aws.GetAmiFilter{
					Name: "virtualization-type",
					Values: []string{
						"hvm",
					},
				},
			},
			MostRecent: &opt0,
			NameRegex:  &opt1,
			Owners: []string{
				"self",
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

