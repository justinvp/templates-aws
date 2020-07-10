package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewInstanceGroup(ctx, "task", &emr.InstanceGroupArgs{
			ClusterId:     pulumi.Any(aws_emr_cluster.Tf - test - cluster.Id),
			InstanceCount: pulumi.Int(1),
			InstanceType:  pulumi.String("m5.xlarge"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

