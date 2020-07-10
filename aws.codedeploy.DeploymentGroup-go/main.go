package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codedeploy"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"codedeploy.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "aWSCodeDeployRole", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole"),
			Role:      exampleRole.Name,
		})
		if err != nil {
			return err
		}
		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", nil)
		if err != nil {
			return err
		}
		exampleTopic, err := sns.NewTopic(ctx, "exampleTopic", nil)
		if err != nil {
			return err
		}
		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
			AlarmConfiguration: &codedeploy.DeploymentGroupAlarmConfigurationArgs{
				Alarms: pulumi.StringArray{
					pulumi.String("my-alarm-name"),
				},
				Enabled: pulumi.Bool(true),
			},
			AppName: exampleApplication.Name,
			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
				Enabled: pulumi.Bool(true),
				Events: pulumi.StringArray{
					pulumi.String("DEPLOYMENT_FAILURE"),
				},
			},
			DeploymentGroupName: pulumi.String("example-group"),
			Ec2TagSets: codedeploy.DeploymentGroupEc2TagSetArray{
				&codedeploy.DeploymentGroupEc2TagSetArgs{
					Ec2TagFilter: pulumi.StringMapArray{
						pulumi.StringMap{
							"key":   pulumi.String("filterkey1"),
							"type":  pulumi.String("KEY_AND_VALUE"),
							"value": pulumi.String("filtervalue"),
						},
						pulumi.StringMap{
							"key":   pulumi.String("filterkey2"),
							"type":  pulumi.String("KEY_AND_VALUE"),
							"value": pulumi.String("filtervalue"),
						},
					},
				},
			},
			ServiceRoleArn: exampleRole.Arn,
			TriggerConfigurations: codedeploy.DeploymentGroupTriggerConfigurationArray{
				&codedeploy.DeploymentGroupTriggerConfigurationArgs{
					TriggerEvents: pulumi.StringArray{
						pulumi.String("DeploymentFailure"),
					},
					TriggerName:      pulumi.String("example-trigger"),
					TriggerTargetArn: exampleTopic.Arn,
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

