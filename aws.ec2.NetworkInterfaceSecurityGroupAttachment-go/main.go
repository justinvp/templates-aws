package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := true
		ami, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
			Filters: []aws.GetAmiFilter{
				aws.GetAmiFilter{
					Name: "name",
					Values: []string{
						"amzn-ami-hvm-*",
					},
				},
			},
			MostRecent: &opt0,
			Owners: []string{
				"amazon",
			},
		}, nil)
		if err != nil {
			return err
		}
		instance, err := ec2.NewInstance(ctx, "instance", &ec2.InstanceArgs{
			Ami:          pulumi.String(ami.Id),
			InstanceType: pulumi.String("t2.micro"),
			Tags: pulumi.StringMap{
				"type": pulumi.String("test-instance"),
			},
		})
		if err != nil {
			return err
		}
		sg, err := ec2.NewSecurityGroup(ctx, "sg", &ec2.SecurityGroupArgs{
			Tags: pulumi.StringMap{
				"type": pulumi.String("test-security-group"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkInterfaceSecurityGroupAttachment(ctx, "sgAttachment", &ec2.NetworkInterfaceSecurityGroupAttachmentArgs{
			NetworkInterfaceId: instance.PrimaryNetworkInterfaceId,
			SecurityGroupId:    sg.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

