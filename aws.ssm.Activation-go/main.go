package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v", "  {\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\"Service\": \"ssm.amazonaws.com\"},\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  }\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "testAttach", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"),
			Role:      testRole.Name,
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewActivation(ctx, "foo", &ssm.ActivationArgs{
			Description:       pulumi.String("Test"),
			IamRole:           testRole.ID(),
			RegistrationLimit: pulumi.Int(5),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_iam_role_policy_attachment.test_attach",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

