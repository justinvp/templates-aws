package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testSnapshotCopyGrant, err := redshift.NewSnapshotCopyGrant(ctx, "testSnapshotCopyGrant", &redshift.SnapshotCopyGrantArgs{
			SnapshotCopyGrantName: pulumi.String("my-grant"),
		})
		if err != nil {
			return err
		}
		_, err = redshift.NewCluster(ctx, "testCluster", &redshift.ClusterArgs{
			SnapshotCopy: &redshift.ClusterSnapshotCopyArgs{
				DestinationRegion: pulumi.String("us-east-2"),
				GrantName:         testSnapshotCopyGrant.SnapshotCopyGrantName,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

