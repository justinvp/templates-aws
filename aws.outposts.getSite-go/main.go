package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/outposts"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "example"
		_, err := outposts.GetSite(ctx, &outposts.GetSiteArgs{
			Name: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

