package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.LookupStateMachine(ctx, &sfn.LookupStateMachineArgs{
			Name: "an_example_sfn_name",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

