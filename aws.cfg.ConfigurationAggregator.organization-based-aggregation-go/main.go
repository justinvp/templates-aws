package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		organizationRole, err := iam.NewRole(ctx, "organizationRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"config.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewConfigurationAggregator(ctx, "organizationConfigurationAggregator", &cfg.ConfigurationAggregatorArgs{
			OrganizationAggregationSource: &cfg.ConfigurationAggregatorOrganizationAggregationSourceArgs{
				AllRegions: pulumi.Bool(true),
				RoleArn:    organizationRole.Arn,
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_iam_role_policy_attachment.organization",
		}))
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "organizationRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations"),
			Role:      organizationRole.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

