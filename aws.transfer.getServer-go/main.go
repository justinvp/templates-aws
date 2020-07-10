package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.LookupServer(ctx, &transfer.LookupServerArgs{
			ServerId: "s-1234567",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

