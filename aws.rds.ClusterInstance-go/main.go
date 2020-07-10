package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "_default", &rds.ClusterArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			ClusterIdentifier: pulumi.String("aurora-cluster-demo"),
			DatabaseName:      pulumi.String("mydb"),
			MasterPassword:    pulumi.String("barbut8chars"),
			MasterUsername:    pulumi.String("foo"),
		})
		if err != nil {
			return err
		}
		var clusterInstances []*rds.ClusterInstance
		for key0, val0 := range 2 {
			__res, err := rds.NewClusterInstance(ctx, fmt.Sprintf("clusterInstances-%v", key0), &rds.ClusterInstanceArgs{
				ClusterIdentifier: _default.ID(),
				Identifier:        pulumi.String(fmt.Sprintf("%v%v", "aurora-cluster-demo-", val0)),
				InstanceClass:     pulumi.String("db.r4.large"),
			})
			if err != nil {
				return err
			}
			clusterInstances = append(clusterInstances, __res)
		}
		return nil
	})
}

