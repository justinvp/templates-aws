package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bar, err := rds.NewInstance(ctx, "bar", &rds.InstanceArgs{
			AllocatedStorage:      pulumi.Int(10),
			BackupRetentionPeriod: pulumi.Int(0),
			Engine:                pulumi.String("MySQL"),
			EngineVersion:         pulumi.String("5.6.21"),
			InstanceClass:         pulumi.String("db.t2.micro"),
			MaintenanceWindow:     pulumi.String("Fri:09:00-Fri:09:30"),
			Name:                  pulumi.String("baz"),
			ParameterGroupName:    pulumi.String("default.mysql5.6"),
			Password:              pulumi.String("barbarbarbar"),
			Username:              pulumi.String("foo"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewSnapshot(ctx, "test", &rds.SnapshotArgs{
			DbInstanceIdentifier: bar.ID(),
			DbSnapshotIdentifier: pulumi.String("testsnapshot1234"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

