package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.LookupService(ctx, &ecs.LookupServiceArgs{
			ClusterArn:  data.Aws_ecs_cluster.Example.Arn,
			ServiceName: "example",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

