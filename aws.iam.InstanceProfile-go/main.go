package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Action\": \"sts:AssumeRole\",\n", "            \"Principal\": {\n", "               \"Service\": \"ec2.amazonaws.com\"\n", "            },\n", "            \"Effect\": \"Allow\",\n", "            \"Sid\": \"\"\n", "        }\n", "    ]\n", "}\n", "\n")),
			Path:             pulumi.String("/"),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewInstanceProfile(ctx, "testProfile", &iam.InstanceProfileArgs{
			Role: role.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

