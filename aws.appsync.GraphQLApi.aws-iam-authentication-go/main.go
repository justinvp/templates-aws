package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("AWS_IAM"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

