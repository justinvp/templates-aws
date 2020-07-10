package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewCluster(ctx, "_default", &redshift.ClusterArgs{
			ClusterIdentifier: pulumi.String("tf-redshift-cluster"),
			ClusterType:       pulumi.String("single-node"),
			DatabaseName:      pulumi.String("mydb"),
			MasterPassword:    pulumi.String("Mustbe8characters"),
			MasterUsername:    pulumi.String("foo"),
			NodeType:          pulumi.String("dc1.large"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

