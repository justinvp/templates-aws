package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewAgent(ctx, "example", &datasync.AgentArgs{
			IpAddress: pulumi.String("1.2.3.4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

