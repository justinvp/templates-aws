package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "    {\n", "      \"Version\": \"2012-10-17\",\n", "      \"Statement\": [\n", "        {\n", "          \"Action\": \"sts:AssumeRole\",\n", "          \"Principal\": {\n", "            \"Service\": \"ec2.amazonaws.com\"\n", "          },\n", "          \"Effect\": \"Allow\",\n", "          \"Sid\": \"\"\n", "        }\n", "      ]\n", "    }\n", "\n")),
		})
		if err != nil {
			return err
		}
		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
			Description: pulumi.String("A test policy"),
			Policy:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"ec2:Describe*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "test_attach", &iam.RolePolicyAttachmentArgs{
			PolicyArn: policy.Arn,
			Role:      role.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

