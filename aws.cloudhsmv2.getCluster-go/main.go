package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudhsmv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudhsmv2.LookupCluster(ctx, &cloudhsmv2.LookupClusterArgs{
			ClusterId: "cluster-testclusterid",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

