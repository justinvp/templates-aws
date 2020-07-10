package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/simpledb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := simpledb.NewDomain(ctx, "users", nil)
		if err != nil {
			return err
		}
		return nil
	})
}

