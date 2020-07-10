package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/msk"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := msk.LookupConfiguration(ctx, &msk.LookupConfigurationArgs{
			Name: "example",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

