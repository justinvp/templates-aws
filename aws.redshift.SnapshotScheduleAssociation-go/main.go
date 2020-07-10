package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultCluster, err := redshift.NewCluster(ctx, "defaultCluster", &redshift.ClusterArgs{
			ClusterIdentifier: pulumi.String("tf-redshift-cluster"),
			ClusterType:       pulumi.String("single-node"),
			DatabaseName:      pulumi.String("mydb"),
			MasterPassword:    pulumi.String("Mustbe8characters"),
			MasterUsername:    pulumi.String("foo"),
			NodeType:          pulumi.String("dc1.large"),
		})
		if err != nil {
			return err
		}
		defaultSnapshotSchedule, err := redshift.NewSnapshotSchedule(ctx, "defaultSnapshotSchedule", &redshift.SnapshotScheduleArgs{
			Definitions: pulumi.StringArray{
				pulumi.String("rate(12 hours)"),
			},
			Identifier: pulumi.String("tf-redshift-snapshot-schedule"),
		})
		if err != nil {
			return err
		}
		_, err = redshift.NewSnapshotScheduleAssociation(ctx, "defaultSnapshotScheduleAssociation", &redshift.SnapshotScheduleAssociationArgs{
			ClusterIdentifier:  defaultCluster.ID(),
			ScheduleIdentifier: defaultSnapshotSchedule.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

