package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dlm"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dlmLifecycleRole, err := iam.NewRole(ctx, "dlmLifecycleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"dlm.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "dlmLifecycle", &iam.RolePolicyArgs{
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "   \"Version\": \"2012-10-17\",\n", "   \"Statement\": [\n", "      {\n", "         \"Effect\": \"Allow\",\n", "         \"Action\": [\n", "            \"ec2:CreateSnapshot\",\n", "            \"ec2:DeleteSnapshot\",\n", "            \"ec2:DescribeVolumes\",\n", "            \"ec2:DescribeSnapshots\"\n", "         ],\n", "         \"Resource\": \"*\"\n", "      },\n", "      {\n", "         \"Effect\": \"Allow\",\n", "         \"Action\": [\n", "            \"ec2:CreateTags\"\n", "         ],\n", "         \"Resource\": \"arn:aws:ec2:*::snapshot/*\"\n", "      }\n", "   ]\n", "}\n", "\n")),
			Role:   dlmLifecycleRole.ID(),
		})
		if err != nil {
			return err
		}
		_, err = dlm.NewLifecyclePolicy(ctx, "example", &dlm.LifecyclePolicyArgs{
			Description:      pulumi.String("example DLM lifecycle policy"),
			ExecutionRoleArn: dlmLifecycleRole.Arn,
			PolicyDetails: &dlm.LifecyclePolicyPolicyDetailsArgs{
				ResourceTypes: pulumi.StringArray{
					pulumi.String("VOLUME"),
				},
				Schedule: pulumi.MapArray{
					pulumi.Map{
						"copyTags": pulumi.Bool(false),
						"createRule": pulumi.Map{
							"interval":     pulumi.Float64(24),
							"intervalUnit": pulumi.String("HOURS"),
							"times":        pulumi.String("23:45"),
						},
						"name": pulumi.String("2 weeks of daily snapshots"),
						"retainRule": pulumi.Float64Map{
							"count": pulumi.Float64(14),
						},
						"tagsToAdd": pulumi.StringMap{
							"SnapshotCreator": pulumi.String("DLM"),
						},
					},
				},
				TargetTags: pulumi.StringMap{
					"Snapshot": pulumi.String("true"),
				},
			},
			State: pulumi.String("ENABLED"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

