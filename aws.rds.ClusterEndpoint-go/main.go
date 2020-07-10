package main

import (
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
			BackupRetentionPeriod: pulumi.Int(5),
			ClusterIdentifier:     pulumi.String("aurora-cluster-demo"),
			DatabaseName:          pulumi.String("mydb"),
			MasterPassword:        pulumi.String("bar"),
			MasterUsername:        pulumi.String("foo"),
			PreferredBackupWindow: pulumi.String("07:00-09:00"),
		})
		if err != nil {
			return err
		}
		test1, err := rds.NewClusterInstance(ctx, "test1", &rds.ClusterInstanceArgs{
			ApplyImmediately:  pulumi.Bool(true),
			ClusterIdentifier: _default.ID(),
			Identifier:        pulumi.String("test1"),
			InstanceClass:     pulumi.String("db.t2.small"),
		})
		if err != nil {
			return err
		}
		test2, err := rds.NewClusterInstance(ctx, "test2", &rds.ClusterInstanceArgs{
			ApplyImmediately:  pulumi.Bool(true),
			ClusterIdentifier: _default.ID(),
			Identifier:        pulumi.String("test2"),
			InstanceClass:     pulumi.String("db.t2.small"),
		})
		if err != nil {
			return err
		}
		test3, err := rds.NewClusterInstance(ctx, "test3", &rds.ClusterInstanceArgs{
			ApplyImmediately:  pulumi.Bool(true),
			ClusterIdentifier: _default.ID(),
			Identifier:        pulumi.String("test3"),
			InstanceClass:     pulumi.String("db.t2.small"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewClusterEndpoint(ctx, "eligible", &rds.ClusterEndpointArgs{
			ClusterEndpointIdentifier: pulumi.String("reader"),
			ClusterIdentifier:         _default.ID(),
			CustomEndpointType:        pulumi.String("READER"),
			ExcludedMembers: pulumi.StringArray{
				test1.ID(),
				test2.ID(),
			},
		})
		if err != nil {
			return err
		}
		_, err = rds.NewClusterEndpoint(ctx, "static", &rds.ClusterEndpointArgs{
			ClusterEndpointIdentifier: pulumi.String("static"),
			ClusterIdentifier:         _default.ID(),
			CustomEndpointType:        pulumi.String("READER"),
			StaticMembers: pulumi.StringArray{
				test1.ID(),
				test3.ID(),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

