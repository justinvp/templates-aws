package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mainIdentityPool, err := cognito.NewIdentityPool(ctx, "mainIdentityPool", &cognito.IdentityPoolArgs{
			AllowUnauthenticatedIdentities: pulumi.Bool(false),
			IdentityPoolName:               pulumi.String("identity pool"),
			SupportedLoginProviders: pulumi.StringMap{
				"graph.facebook.com": pulumi.String("7346241598935555"),
			},
		})
		if err != nil {
			return err
		}
		authenticatedRole, err := iam.NewRole(ctx, "authenticatedRole", &iam.RoleArgs{
			AssumeRolePolicy: mainIdentityPool.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Federated\": \"cognito-identity.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRoleWithWebIdentity\",\n", "      \"Condition\": {\n", "        \"StringEquals\": {\n", "          \"cognito-identity.amazonaws.com:aud\": \"", id, "\"\n", "        },\n", "        \"ForAnyValue:StringLike\": {\n", "          \"cognito-identity.amazonaws.com:amr\": \"authenticated\"\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "authenticatedRolePolicy", &iam.RolePolicyArgs{
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Action\": [\n", "        \"mobileanalytics:PutEvents\",\n", "        \"cognito-sync:*\",\n", "        \"cognito-identity:*\"\n", "      ],\n", "      \"Resource\": [\n", "        \"*\"\n", "      ]\n", "    }\n", "  ]\n", "}\n", "\n")),
			Role:   authenticatedRole.ID(),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewIdentityPoolRoleAttachment(ctx, "mainIdentityPoolRoleAttachment", &cognito.IdentityPoolRoleAttachmentArgs{
			IdentityPoolId: mainIdentityPool.ID(),
			RoleMappings: cognito.IdentityPoolRoleAttachmentRoleMappingArray{
				&cognito.IdentityPoolRoleAttachmentRoleMappingArgs{
					AmbiguousRoleResolution: pulumi.String("AuthenticatedRole"),
					IdentityProvider:        pulumi.String("graph.facebook.com"),
					MappingRule: pulumi.StringMapArray{
						pulumi.StringMap{
							"claim":     pulumi.String("isAdmin"),
							"matchType": pulumi.String("Equals"),
							"roleArn":   authenticatedRole.Arn,
							"value":     pulumi.String("paid"),
						},
					},
					Type: pulumi.String("Rules"),
				},
			},
			Roles: pulumi.StringMap{
				"authenticated": authenticatedRole.Arn,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

