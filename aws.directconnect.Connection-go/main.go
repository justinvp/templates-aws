package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.NewConnection(ctx, "hoge", &directconnect.ConnectionArgs{
			Bandwidth: pulumi.String("1Gbps"),
			Location:  pulumi.String("EqDC2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

