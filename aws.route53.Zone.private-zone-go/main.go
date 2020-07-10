package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewZone(ctx, "private", &route53.ZoneArgs{
			Vpcs: route53.ZoneVpcArray{
				&route53.ZoneVpcArgs{
					VpcId: pulumi.Any(aws_vpc.Example.Id),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

