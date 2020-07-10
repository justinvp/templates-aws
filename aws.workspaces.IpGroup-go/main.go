package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workspaces.NewIpGroup(ctx, "contractors", &workspaces.IpGroupArgs{
			Description: pulumi.String("Contractors IP access control group"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

