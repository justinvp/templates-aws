package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewUser(ctx, "example", &quicksight.UserArgs{
			Email:        pulumi.String("author@example.com"),
			IdentityType: pulumi.String("IAM"),
			UserName:     pulumi.String("an-author"),
			UserRole:     pulumi.String("AUTHOR"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

