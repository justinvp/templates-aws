package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooDeploymentConfig, err := codedeploy.NewDeploymentConfig(ctx, "fooDeploymentConfig", &codedeploy.DeploymentConfigArgs{
			DeploymentConfigName: pulumi.String("test-deployment-config"),
			MinimumHealthyHosts: &codedeploy.DeploymentConfigMinimumHealthyHostsArgs{
				Type:  pulumi.String("HOST_COUNT"),
				Value: pulumi.Int(2),
			},
		})
		if err != nil {
			return err
		}
		_, err = codedeploy.NewDeploymentGroup(ctx, "fooDeploymentGroup", &codedeploy.DeploymentGroupArgs{
			AlarmConfiguration: &codedeploy.DeploymentGroupAlarmConfigurationArgs{
				Alarms: pulumi.StringArray{
					pulumi.String("my-alarm-name"),
				},
				Enabled: pulumi.Bool(true),
			},
			AppName: pulumi.Any(aws_codedeploy_app.Foo_app.Name),
			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
				Enabled: pulumi.Bool(true),
				Events: pulumi.StringArray{
					pulumi.String("DEPLOYMENT_FAILURE"),
				},
			},
			DeploymentConfigName: fooDeploymentConfig.ID(),
			DeploymentGroupName:  pulumi.String("bar"),
			Ec2TagFilters: codedeploy.DeploymentGroupEc2TagFilterArray{
				&codedeploy.DeploymentGroupEc2TagFilterArgs{
					Key:   pulumi.String("filterkey"),
					Type:  pulumi.String("KEY_AND_VALUE"),
					Value: pulumi.String("filtervalue"),
				},
			},
			ServiceRoleArn: pulumi.Any(aws_iam_role.Foo_role.Arn),
			TriggerConfigurations: codedeploy.DeploymentGroupTriggerConfigurationArray{
				&codedeploy.DeploymentGroupTriggerConfigurationArgs{
					TriggerEvents: pulumi.StringArray{
						pulumi.String("DeploymentFailure"),
					},
					TriggerName:      pulumi.String("foo-trigger"),
					TriggerTargetArn: pulumi.String("foo-topic-arn"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

