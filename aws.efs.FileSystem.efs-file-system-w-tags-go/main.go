package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := efs.NewFileSystem(ctx, "foo", &efs.FileSystemArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("MyProduct"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

