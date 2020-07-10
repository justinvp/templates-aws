package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewSubnetGroup(ctx, "_default", &rds.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				pulumi.Any(aws_subnet.Frontend.Id),
				pulumi.Any(aws_subnet.Backend.Id),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("My DB subnet group"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

