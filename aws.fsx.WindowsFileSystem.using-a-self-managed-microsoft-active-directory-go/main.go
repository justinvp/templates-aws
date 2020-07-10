package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewWindowsFileSystem(ctx, "example", &fsx.WindowsFileSystemArgs{
			KmsKeyId: pulumi.Any(aws_kms_key.Example.Arn),
			SelfManagedActiveDirectory: &fsx.WindowsFileSystemSelfManagedActiveDirectoryArgs{
				DnsIps: pulumi.StringArray{
					pulumi.String("10.0.0.111"),
					pulumi.String("10.0.0.222"),
				},
				DomainName: pulumi.String("corp.example.com"),
				Password:   pulumi.String("avoid-plaintext-passwords"),
				Username:   pulumi.String("Admin"),
			},
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

