package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		lbUser, err := iam.NewUser(ctx, "lbUser", &iam.UserArgs{
			Path: pulumi.String("/system/"),
		})
		if err != nil {
			return err
		}
		lbAccessKey, err := iam.NewAccessKey(ctx, "lbAccessKey", &iam.AccessKeyArgs{
			PgpKey: pulumi.String("keybase:some_person_that_exists"),
			User:   lbUser.Name,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewUserPolicy(ctx, "lbRo", &iam.UserPolicyArgs{
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"ec2:Describe*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n", "\n")),
			User:   lbUser.Name,
		})
		if err != nil {
			return err
		}
		ctx.Export("secret", lbAccessKey.EncryptedSecret)
		return nil
	})
}

