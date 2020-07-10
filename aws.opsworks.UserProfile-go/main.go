package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewUserProfile(ctx, "myProfile", &opsworks.UserProfileArgs{
			SshUsername: pulumi.String("my_user"),
			UserArn:     pulumi.Any(aws_iam_user.User.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

