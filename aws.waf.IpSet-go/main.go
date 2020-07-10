package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewIpSet(ctx, "ipset", &waf.IpSetArgs{
			IpSetDescriptors: waf.IpSetIpSetDescriptorArray{
				&waf.IpSetIpSetDescriptorArgs{
					Type:  pulumi.String("IPV4"),
					Value: pulumi.String("192.0.7.0/24"),
				},
				&waf.IpSetIpSetDescriptorArgs{
					Type:  pulumi.String("IPV4"),
					Value: pulumi.String("10.16.16.0/16"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

