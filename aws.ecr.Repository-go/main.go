package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecr.NewRepository(ctx, "foo", &ecr.RepositoryArgs{
			ImageScanningConfiguration: &ecr.RepositoryImageScanningConfigurationArgs{
				ScanOnPush: pulumi.Bool(true),
			},
			ImageTagMutability: pulumi.String("MUTABLE"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

