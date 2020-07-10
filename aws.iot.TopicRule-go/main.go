package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iot"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mytopic, err := sns.NewTopic(ctx, "mytopic", nil)
		if err != nil {
			return err
		}
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"iot.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iot.NewTopicRule(ctx, "rule", &iot.TopicRuleArgs{
			Description: pulumi.String("Example rule"),
			Enabled:     pulumi.Bool(true),
			Sns: &iot.TopicRuleSnsArgs{
				Sns: pulumi.String("RAW"),
				Sns: role.Arn,
				Sns: mytopic.Arn,
			},
			Sql:        pulumi.String("SELECT * FROM 'topic/test'"),
			SqlVersion: pulumi.String("2016-03-23"),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "iamPolicyForLambda", &iam.RolePolicyArgs{
			Policy: mytopic.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "        \"Effect\": \"Allow\",\n", "        \"Action\": [\n", "            \"sns:Publish\"\n", "        ],\n", "        \"Resource\": \"", arn, "\"\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
			Role: role.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

