package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewCluster(ctx, "replica", &elasticache.ClusterArgs{
			ReplicationGroupId: pulumi.Any(aws_elasticache_replication_group.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

