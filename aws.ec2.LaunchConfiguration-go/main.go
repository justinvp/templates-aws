package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := true
		ubuntu, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
			Filters: []aws.GetAmiFilter{
				aws.GetAmiFilter{
					Name: "name",
					Values: []string{
						"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
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
			Owners: []string{
				"099720109477",
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewLaunchConfiguration(ctx, "asConf", &ec2.LaunchConfigurationArgs{
			ImageId:      pulumi.String(ubuntu.Id),
			InstanceType: pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

