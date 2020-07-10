package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dax"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dax.NewCluster(ctx, "bar", &dax.ClusterArgs{
			ClusterName:       pulumi.String("cluster-example"),
			IamRoleArn:        pulumi.Any(data.Aws_iam_role.Example.Arn),
			NodeType:          pulumi.String("dax.r4.large"),
			ReplicationFactor: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

