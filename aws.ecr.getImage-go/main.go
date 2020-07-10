package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "latest"
		_, err := ecr.GetImage(ctx, &ecr.GetImageArgs{
			ImageTag:       &opt0,
			RepositoryName: "my/service",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

