package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "example", &rds.ClusterArgs{
			ClusterIdentifier: pulumi.String("example"),
			DbSubnetGroupName: pulumi.Any(aws_db_subnet_group.Example.Name),
			EngineMode:        pulumi.String("multimaster"),
			MasterPassword:    pulumi.String("barbarbarbar"),
			MasterUsername:    pulumi.String("foo"),
			SkipFinalSnapshot: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

