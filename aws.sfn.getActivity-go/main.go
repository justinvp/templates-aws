package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "my-activity"
		_, err := sfn.LookupActivity(ctx, &sfn.LookupActivityArgs{
			Name: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

