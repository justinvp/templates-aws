package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codecommit"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codecommit.NewRepository(ctx, "test", &codecommit.RepositoryArgs{
			Description:    pulumi.String("This is the Sample App Repository"),
			RepositoryName: pulumi.String("MyTestRepository"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

