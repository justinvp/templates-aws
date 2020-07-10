package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/qldb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := qldb.LookupLedger(ctx, &qldb.LookupLedgerArgs{
			Name: "an_example_ledger",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

