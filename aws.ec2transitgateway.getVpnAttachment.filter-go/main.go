package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.GetVpnAttachment(ctx, &ec2transitgateway.GetVpnAttachmentArgs{
			Filters: []ec2transitgateway.GetVpnAttachmentFilter{
				ec2transitgateway.GetVpnAttachmentFilter{
					Name: "resource-id",
					Values: []string{
						"some-resource",
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

