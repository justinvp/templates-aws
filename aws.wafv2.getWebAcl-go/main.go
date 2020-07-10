package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafv2.GetWebAcl(ctx, &wafv2.GetWebAclArgs{
			Name:  "some-web-acl",
			Scope: "REGIONAL",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

