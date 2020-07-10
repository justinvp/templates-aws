package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := docdb.NewCluster(ctx, "_default", &docdb.ClusterArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			ClusterIdentifier: pulumi.String("docdb-cluster-demo"),
			MasterPassword:    pulumi.String("barbut8chars"),
			MasterUsername:    pulumi.String("foo"),
		})
		if err != nil {
			return err
		}
		var clusterInstances []*docdb.ClusterInstance
		for key0, val0 := range 2 {
			__res, err := docdb.NewClusterInstance(ctx, fmt.Sprintf("clusterInstances-%v", key0), &docdb.ClusterInstanceArgs{
				ClusterIdentifier: _default.ID(),
				Identifier:        pulumi.String(fmt.Sprintf("%v%v", "docdb-cluster-demo-", val0)),
				InstanceClass:     pulumi.String("db.r5.large"),
			})
			if err != nil {
				return err
			}
			clusterInstances = append(clusterInstances, __res)
		}
		return nil
	})
}

