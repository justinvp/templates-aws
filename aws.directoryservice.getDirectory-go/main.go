package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directoryservice"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directoryservice.LookupDirectory(ctx, &directoryservice.LookupDirectoryArgs{
			DirectoryId: aws_directory_service_directory.Main.Id,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

