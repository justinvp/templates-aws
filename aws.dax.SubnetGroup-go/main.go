package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dax"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dax.NewSubnetGroup(ctx, "example", &dax.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				pulumi.Any(aws_subnet.Example1.Id),
				pulumi.Any(aws_subnet.Example2.Id),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

