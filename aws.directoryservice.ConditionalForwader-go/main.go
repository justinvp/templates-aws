package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directoryservice"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directoryservice.NewConditionalForwader(ctx, "example", &directoryservice.ConditionalForwaderArgs{
			DirectoryId: pulumi.Any(aws_directory_service_directory.Ad.Id),
			DnsIps: pulumi.StringArray{
				pulumi.String("8.8.8.8"),
				pulumi.String("8.8.4.4"),
			},
			RemoteDomainName: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

