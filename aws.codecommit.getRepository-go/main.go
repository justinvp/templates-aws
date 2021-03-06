package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codecommit"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codecommit.LookupRepository(ctx, &codecommit.LookupRepositoryArgs{
			RepositoryName: "MyTestRepository",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

