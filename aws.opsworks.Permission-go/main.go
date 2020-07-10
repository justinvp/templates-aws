package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewPermission(ctx, "myStackPermission", &opsworks.PermissionArgs{
			AllowSsh:  pulumi.Bool(true),
			AllowSudo: pulumi.Bool(true),
			Level:     pulumi.String("iam_only"),
			StackId:   pulumi.Any(aws_opsworks_stack.Stack.Id),
			UserArn:   pulumi.Any(aws_iam_user.User.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

