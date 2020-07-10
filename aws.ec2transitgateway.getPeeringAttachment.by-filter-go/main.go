package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.LookupPeeringAttachment(ctx, &ec2transitgateway.LookupPeeringAttachmentArgs{
			Filters: []ec2transitgateway.GetPeeringAttachmentFilter{
				ec2transitgateway.GetPeeringAttachmentFilter{
					Name: "transit-gateway-attachment-id",
					Values: []string{
						"tgw-attach-12345678",
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

