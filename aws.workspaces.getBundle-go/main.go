package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "Value with Windows 10 and Office 2016"
		opt1 := "AMAZON"
		_, err := workspaces.GetBundle(ctx, &workspaces.GetBundleArgs{
			Name:  &opt0,
			Owner: &opt1,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

