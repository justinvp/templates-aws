package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewConnection(ctx, "example", &glue.ConnectionArgs{
			ConnectionProperties: pulumi.StringMap{
				"JDBC_CONNECTION_URL": pulumi.String("jdbc:mysql://example.com/exampledatabase"),
				"PASSWORD":            pulumi.String("examplepassword"),
				"USERNAME":            pulumi.String("exampleusername"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

