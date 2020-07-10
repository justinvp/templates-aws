package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewReplicationSubnetGroup(ctx, "test", &dms.ReplicationSubnetGroupArgs{
			ReplicationSubnetGroupDescription: pulumi.String("Test replication subnet group"),
			ReplicationSubnetGroupId:          pulumi.String("test-dms-replication-subnet-group-tf"),
			SubnetIds: pulumi.StringArray{
				pulumi.String("subnet-12345678"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

