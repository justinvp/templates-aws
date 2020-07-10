package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudhsmv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cluster, err := cloudhsmv2.LookupCluster(ctx, &cloudhsmv2.LookupClusterArgs{
			ClusterId: _var.Cloudhsm_cluster_id,
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudhsmv2.NewHsm(ctx, "cloudhsmV2Hsm", &cloudhsmv2.HsmArgs{
			ClusterId: pulumi.String(cluster.ClusterId),
			SubnetId:  pulumi.String(cluster.SubnetIds[0]),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

