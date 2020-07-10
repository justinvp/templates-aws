package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewReplicationGroup(ctx, "baz", &elasticache.ReplicationGroupArgs{
			AutomaticFailoverEnabled: pulumi.Bool(true),
			ClusterMode: &elasticache.ReplicationGroupClusterModeArgs{
				NumNodeGroups:        pulumi.Int(2),
				ReplicasPerNodeGroup: pulumi.Int(1),
			},
			NodeType:                    pulumi.String("cache.t2.small"),
			ParameterGroupName:          pulumi.String("default.redis3.2.cluster.on"),
			Port:                        pulumi.Int(6379),
			ReplicationGroupDescription: pulumi.String("test description"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

