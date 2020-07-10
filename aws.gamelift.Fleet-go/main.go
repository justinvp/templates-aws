package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/gamelift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gamelift.NewFleet(ctx, "example", &gamelift.FleetArgs{
			BuildId:         pulumi.Any(aws_gamelift_build.Example.Id),
			Ec2InstanceType: pulumi.String("t2.micro"),
			FleetType:       pulumi.String("ON_DEMAND"),
			RuntimeConfiguration: &gamelift.FleetRuntimeConfigurationArgs{
				ServerProcess: pulumi.MapArray{
					pulumi.Map{
						"concurrentExecutions": pulumi.Float64(1),
						"launchPath":           pulumi.String("C:\\game\\GomokuServer.exe"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

