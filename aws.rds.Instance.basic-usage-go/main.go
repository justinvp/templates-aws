package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewInstance(ctx, "_default", &rds.InstanceArgs{
			AllocatedStorage:   pulumi.Int(20),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("5.7"),
			InstanceClass:      pulumi.String("db.t2.micro"),
			Name:               pulumi.String("mydb"),
			ParameterGroupName: pulumi.String("default.mysql5.7"),
			Password:           pulumi.String("foobarbaz"),
			StorageType:        pulumi.String("gp2"),
			Username:           pulumi.String("foo"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

