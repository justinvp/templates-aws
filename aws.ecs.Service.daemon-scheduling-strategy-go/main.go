package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.NewService(ctx, "bar", &ecs.ServiceArgs{
			Cluster:            pulumi.Any(aws_ecs_cluster.Foo.Id),
			SchedulingStrategy: pulumi.String("DAEMON"),
			TaskDefinition:     pulumi.Any(aws_ecs_task_definition.Bar.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

