package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewCluster(ctx, "example", &elasticache.ClusterArgs{
			Engine:             pulumi.String("memcached"),
			NodeType:           pulumi.String("cache.m4.large"),
			NumCacheNodes:      pulumi.Int(2),
			ParameterGroupName: pulumi.String("default.memcached1.4"),
			Port:               pulumi.Int(11211),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

