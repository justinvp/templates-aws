package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := fileSystemId
		_, err := efs.LookupFileSystem(ctx, &efs.LookupFileSystemArgs{
			FileSystemId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

