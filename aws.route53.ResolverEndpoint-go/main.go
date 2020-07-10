package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewResolverEndpoint(ctx, "foo", &route53.ResolverEndpointArgs{
			Direction: pulumi.String("INBOUND"),
			IpAddresses: route53.ResolverEndpointIpAddressArray{
				&route53.ResolverEndpointIpAddressArgs{
					SubnetId: pulumi.Any(aws_subnet.Sn1.Id),
				},
				&route53.ResolverEndpointIpAddressArgs{
					Ip:       pulumi.String("10.0.64.4"),
					SubnetId: pulumi.Any(aws_subnet.Sn2.Id),
				},
			},
			SecurityGroupIds: pulumi.StringArray{
				pulumi.Any(aws_security_group.Sg1.Id),
				pulumi.Any(aws_security_group.Sg2.Id),
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Prod"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

