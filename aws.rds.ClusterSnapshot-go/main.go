package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewClusterSnapshot(ctx, "example", &rds.ClusterSnapshotArgs{
			DbClusterIdentifier:         pulumi.Any(aws_rds_cluster.Example.Id),
			DbClusterSnapshotIdentifier: pulumi.String("resourcetestsnapshot1234"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

