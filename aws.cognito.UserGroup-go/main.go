package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mainUserPool, err := cognito.NewUserPool(ctx, "mainUserPool", nil)
		if err != nil {
			return err
		}
		groupRole, err := iam.NewRole(ctx, "groupRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Federated\": \"cognito-identity.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRoleWithWebIdentity\",\n", "      \"Condition\": {\n", "        \"StringEquals\": {\n", "          \"cognito-identity.amazonaws.com:aud\": \"us-east-1:12345678-dead-beef-cafe-123456790ab\"\n", "        },\n", "        \"ForAnyValue:StringLike\": {\n", "          \"cognito-identity.amazonaws.com:amr\": \"authenticated\"\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewUserGroup(ctx, "mainUserGroup", &cognito.UserGroupArgs{
			Description: pulumi.String("Managed by Pulumi"),
			Precedence:  pulumi.Int(42),
			RoleArn:     groupRole.Arn,
			UserPoolId:  mainUserPool.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

