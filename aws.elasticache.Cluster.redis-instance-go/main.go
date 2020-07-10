package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewCluster(ctx, "example", &elasticache.ClusterArgs{
			Engine:             pulumi.String("redis"),
			EngineVersion:      pulumi.String("3.2.10"),
			NodeType:           pulumi.String("cache.m4.large"),
			NumCacheNodes:      pulumi.Int(1),
			ParameterGroupName: pulumi.String("default.redis3.2"),
			Port:               pulumi.Int(6379),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

