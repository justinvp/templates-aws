package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewSecurityGroup(ctx, "_default", &redshift.SecurityGroupArgs{
			Ingress: redshift.SecurityGroupIngressArray{
				&redshift.SecurityGroupIngressArgs{
					Cidr: pulumi.String("10.0.0.0/24"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

