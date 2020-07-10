package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := docdb.NewCluster(ctx, "docdb", &docdb.ClusterArgs{
			BackupRetentionPeriod: pulumi.Int(5),
			ClusterIdentifier:     pulumi.String("my-docdb-cluster"),
			Engine:                pulumi.String("docdb"),
			MasterPassword:        pulumi.String("mustbeeightchars"),
			MasterUsername:        pulumi.String("foo"),
			PreferredBackupWindow: pulumi.String("07:00-09:00"),
			SkipFinalSnapshot:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

