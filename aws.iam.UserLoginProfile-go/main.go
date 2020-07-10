package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleUser, err := iam.NewUser(ctx, "exampleUser", &iam.UserArgs{
			ForceDestroy: pulumi.Bool(true),
			Path:         pulumi.String("/"),
		})
		if err != nil {
			return err
		}
		exampleUserLoginProfile, err := iam.NewUserLoginProfile(ctx, "exampleUserLoginProfile", &iam.UserLoginProfileArgs{
			PgpKey: pulumi.String("keybase:some_person_that_exists"),
			User:   exampleUser.Name,
		})
		if err != nil {
			return err
		}
		ctx.Export("password", exampleUserLoginProfile.EncryptedPassword)
		return nil
	})
}

