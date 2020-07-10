package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/backup"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\"sts:AssumeRole\"],\n", "      \"Effect\": \"allow\",\n", "      \"Principal\": {\n", "        \"Service\": [\"backup.amazonaws.com\"]\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup"),
			Role:      exampleRole.Name,
		})
		if err != nil {
			return err
		}
		_, err = backup.NewSelection(ctx, "exampleSelection", &backup.SelectionArgs{
			IamRoleArn: exampleRole.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

