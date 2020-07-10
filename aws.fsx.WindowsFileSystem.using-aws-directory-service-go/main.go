package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewWindowsFileSystem(ctx, "example", &fsx.WindowsFileSystemArgs{
			ActiveDirectoryId:  pulumi.Any(aws_directory_service_directory.Example.Id),
			KmsKeyId:           pulumi.Any(aws_kms_key.Example.Arn),
			StorageCapacity:    pulumi.Int(300),
			SubnetIds:          pulumi.Any(aws_subnet.Example.Id),
			ThroughputCapacity: pulumi.Int(1024),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

