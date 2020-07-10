package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := efs.LookupAccessPoint(ctx, &efs.LookupAccessPointArgs{
			AccessPointId: "fsap-12345678",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

