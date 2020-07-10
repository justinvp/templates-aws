package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pool, err := cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewResourceServer(ctx, "resource", &cognito.ResourceServerArgs{
			Identifier: pulumi.String("https://example.com"),
			UserPoolId: pool.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

