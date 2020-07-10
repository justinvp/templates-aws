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
			Engine:                pulumi.String("aurora-mysql"),
			EngineVersion:         pulumi.String("5.7.mysql_aurora.2.03.2"),
			MasterPassword:        pulumi.String("bar"),
			MasterUsername:        pulumi.String("foo"),
			PreferredBackupWindow: pulumi.String("07:00-09:00"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

