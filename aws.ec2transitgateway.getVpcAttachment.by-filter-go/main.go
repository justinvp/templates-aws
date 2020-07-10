package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.LookupVpcAttachment(ctx, &ec2transitgateway.LookupVpcAttachmentArgs{
			Filters: []ec2transitgateway.GetVpcAttachmentFilter{
				ec2transitgateway.GetVpcAttachmentFilter{
					Name: "vpc-id",
					Values: []string{
						"vpc-12345678",
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

