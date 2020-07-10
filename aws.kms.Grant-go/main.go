package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		key, err := kms.NewKey(ctx, "key", nil)
		if err != nil {
			return err
		}
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"lambda.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = kms.NewGrant(ctx, "grant", &kms.GrantArgs{
			Constraints: kms.GrantConstraintArray{
				&kms.GrantConstraintArgs{
					EncryptionContextEquals: pulumi.StringMap{
						"Department": pulumi.String("Finance"),
					},
				},
			},
			GranteePrincipal: role.Arn,
			KeyId:            key.KeyId,
			Operations: pulumi.StringArray{
				pulumi.String("Encrypt"),
				pulumi.String("Decrypt"),
				pulumi.String("GenerateDataKey"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

