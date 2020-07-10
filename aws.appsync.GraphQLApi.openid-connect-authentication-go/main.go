package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("OPENID_CONNECT"),
			OpenidConnectConfig: &appsync.GraphQLApiOpenidConnectConfigArgs{
				Issuer: pulumi.String("https://example.com"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

