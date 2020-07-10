package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("API_KEY"),
		})
		if err != nil {
			return err
		}
		_, err = appsync.NewApiKey(ctx, "exampleApiKey", &appsync.ApiKeyArgs{
			ApiId:   exampleGraphQLApi.ID(),
			Expires: pulumi.String("2018-05-03T04:00:00Z"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

