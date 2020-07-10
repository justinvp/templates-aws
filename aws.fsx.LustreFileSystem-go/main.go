package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewLustreFileSystem(ctx, "example", &fsx.LustreFileSystemArgs{
			ImportPath:      pulumi.String(fmt.Sprintf("%v%v", "s3://", aws_s3_bucket.Example.Bucket)),
			StorageCapacity: pulumi.Int(1200),
			SubnetIds:       pulumi.Any(aws_subnet.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

