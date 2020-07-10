package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "AWS-"
		opt1 := "CENTOS"
		_, err := ssm.LookupPatchBaseline(ctx, &ssm.LookupPatchBaselineArgs{
			NamePrefix:      &opt0,
			OperatingSystem: &opt1,
			Owner:           "AWS",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

