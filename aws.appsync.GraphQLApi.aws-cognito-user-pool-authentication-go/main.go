package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("AMAZON_COGNITO_USER_POOLS"),
			UserPoolConfig: &appsync.GraphQLApiUserPoolConfigArgs{
				AwsRegion:     pulumi.Any(data.Aws_region.Current.Name),
				DefaultAction: pulumi.String("DENY"),
				UserPoolId:    pulumi.Any(aws_cognito_user_pool.Example.Id),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

