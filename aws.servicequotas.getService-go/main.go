package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/servicequotas"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicequotas.GetService(ctx, &servicequotas.GetServiceArgs{
			ServiceName: "Amazon Virtual Private Cloud (Amazon VPC)",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

