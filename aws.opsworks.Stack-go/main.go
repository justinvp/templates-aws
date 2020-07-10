package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewStack(ctx, "main", &opsworks.StackArgs{
			CustomJson:                pulumi.String(fmt.Sprintf("%v%v%v%v%v%v", "{\n", " \"foobar\": {\n", "    \"version\": \"1.0.0\"\n", "  }\n", "}\n", "\n")),
			DefaultInstanceProfileArn: pulumi.Any(aws_iam_instance_profile.Opsworks.Arn),
			Region:                    pulumi.String("us-west-1"),
			ServiceRoleArn:            pulumi.Any(aws_iam_role.Opsworks.Arn),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foobar-stack"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

