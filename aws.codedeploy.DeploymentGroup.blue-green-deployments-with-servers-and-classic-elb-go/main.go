package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", nil)
		if err != nil {
			return err
		}
		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
			AppName: exampleApplication.Name,
			BlueGreenDeploymentConfig: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigArgs{
				DeploymentReadyOption: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs{
					ActionOnTimeout:   pulumi.String("STOP_DEPLOYMENT"),
					WaitTimeInMinutes: pulumi.Int(60),
				},
				GreenFleetProvisioningOption: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs{
					Action: pulumi.String("DISCOVER_EXISTING"),
				},
				TerminateBlueInstancesOnDeploymentSuccess: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs{
					Action: pulumi.String("KEEP_ALIVE"),
				},
			},
			DeploymentGroupName: pulumi.String("example-group"),
			DeploymentStyle: &codedeploy.DeploymentGroupDeploymentStyleArgs{
				DeploymentOption: pulumi.String("WITH_TRAFFIC_CONTROL"),
				DeploymentType:   pulumi.String("BLUE_GREEN"),
			},
			LoadBalancerInfo: &codedeploy.DeploymentGroupLoadBalancerInfoArgs{
				ElbInfo: pulumi.AnyMapArray{
					pulumi.AnyMap{
						"name": pulumi.Any(aws_elb.Example.Name),
					},
				},
			},
			ServiceRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

