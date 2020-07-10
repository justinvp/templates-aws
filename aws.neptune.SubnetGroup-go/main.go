package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/neptune"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := neptune.NewSubnetGroup(ctx, "_default", &neptune.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				pulumi.Any(aws_subnet.Frontend.Id),
				pulumi.Any(aws_subnet.Backend.Id),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("My neptune subnet group"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

