package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cfg.NewRule(ctx, "rule", &cfg.RuleArgs{
			Source: &cfg.RuleSourceArgs{
				Owner:            pulumi.String("AWS"),
				SourceIdentifier: pulumi.String("S3_BUCKET_VERSIONING_ENABLED"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_config_configuration_recorder.foo",
		}))
		if err != nil {
			return err
		}
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"config.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewRecorder(ctx, "foo", &cfg.RecorderArgs{
			RoleArn: role.Arn,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "  	{\n", "  		\"Action\": \"config:Put*\",\n", "  		\"Effect\": \"Allow\",\n", "  		\"Resource\": \"*\"\n", "\n", "  	}\n", "  ]\n", "}\n", "\n")),
			Role: role.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

