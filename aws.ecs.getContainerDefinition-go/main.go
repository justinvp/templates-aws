package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.GetContainerDefinition(ctx, &ecs.GetContainerDefinitionArgs{
			ContainerName:  "mongodb",
			TaskDefinition: aws_ecs_task_definition.Mongo.Id,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

