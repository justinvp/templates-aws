package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providers.Newaws(ctx, "primary", &providers.awsArgs{
			Region: pulumi.String("us-east-2"),
		})
		if err != nil {
			return err
		}
		_, err = providers.Newaws(ctx, "secondary", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
			GlobalClusterIdentifier: pulumi.String("example"),
		}, pulumi.Provider("aws.primary"))
		if err != nil {
			return err
		}
		primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
			EngineMode:              pulumi.String("global"),
			GlobalClusterIdentifier: example.ID(),
		}, pulumi.Provider("aws.primary"))
		if err != nil {
			return err
		}
		_, err = rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
			ClusterIdentifier: primaryCluster.ID(),
		}, pulumi.Provider("aws.primary"))
		if err != nil {
			return err
		}
		secondaryCluster, err := rds.NewCluster(ctx, "secondaryCluster", &rds.ClusterArgs{
			EngineMode:              pulumi.String("global"),
			GlobalClusterIdentifier: example.ID(),
		}, pulumi.Provider("aws.secondary"), pulumi.DependsOn([]pulumi.Resource{
			"aws_rds_cluster_instance.primary",
		}))
		if err != nil {
			return err
		}
		_, err = rds.NewClusterInstance(ctx, "secondaryClusterInstance", &rds.ClusterInstanceArgs{
			ClusterIdentifier: secondaryCluster.ID(),
		}, pulumi.Provider("aws.secondary"))
		if err != nil {
			return err
		}
		return nil
	})
}

