package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := docdb.NewClusterSnapshot(ctx, "example", &docdb.ClusterSnapshotArgs{
			DbClusterIdentifier:         pulumi.Any(aws_docdb_cluster.Example.Id),
			DbClusterSnapshotIdentifier: pulumi.String("resourcetestsnapshot1234"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

