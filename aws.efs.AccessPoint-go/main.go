package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := efs.NewAccessPoint(ctx, "test", &efs.AccessPointArgs{
			FileSystemId: pulumi.Any(aws_efs_file_system.Foo.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

