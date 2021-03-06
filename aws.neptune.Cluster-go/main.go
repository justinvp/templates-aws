package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/neptune"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := neptune.NewCluster(ctx, "_default", &neptune.ClusterArgs{
			ApplyImmediately:                 pulumi.Bool(true),
			BackupRetentionPeriod:            pulumi.Int(5),
			ClusterIdentifier:                pulumi.String("neptune-cluster-demo"),
			Engine:                           pulumi.String("neptune"),
			IamDatabaseAuthenticationEnabled: pulumi.Bool(true),
			PreferredBackupWindow:            pulumi.String("07:00-09:00"),
			SkipFinalSnapshot:                pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

