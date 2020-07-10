package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooDeploymentConfig, err := codedeploy.NewDeploymentConfig(ctx, "fooDeploymentConfig", &codedeploy.DeploymentConfigArgs{
			ComputePlatform:      pulumi.String("Lambda"),
			DeploymentConfigName: pulumi.String("test-deployment-config"),
			TrafficRoutingConfig: &codedeploy.DeploymentConfigTrafficRoutingConfigArgs{
				TimeBasedLinear: &codedeploy.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs{
					Interval:   pulumi.Int(10),
					Percentage: pulumi.Int(10),
				},
				Type: pulumi.String("TimeBasedLinear"),
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
					pulumi.String("DEPLOYMENT_STOP_ON_ALARM"),
				},
			},
			DeploymentConfigName: fooDeploymentConfig.ID(),
			DeploymentGroupName:  pulumi.String("bar"),
			ServiceRoleArn:       pulumi.Any(aws_iam_role.Foo_role.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

