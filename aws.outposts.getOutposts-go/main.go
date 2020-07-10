package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/outposts"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := data.Aws_outposts_site.Id
		_, err := outposts.GetOutposts(ctx, &outposts.GetOutpostsArgs{
			SiteId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

