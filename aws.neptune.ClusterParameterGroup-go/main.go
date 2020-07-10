package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/neptune"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := neptune.NewClusterParameterGroup(ctx, "example", &neptune.ClusterParameterGroupArgs{
			Description: pulumi.String("neptune cluster parameter group"),
			Family:      pulumi.String("neptune1"),
			Parameters: neptune.ClusterParameterGroupParameterArray{
				&neptune.ClusterParameterGroupParameterArgs{
					Name:  pulumi.String("neptune_enable_audit_log"),
					Value: pulumi.String("1"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

