package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.LookupCluster(ctx, &ecs.LookupClusterArgs{
			ClusterName: "ecs-mongo-production",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

