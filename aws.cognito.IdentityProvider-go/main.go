package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := cognito.NewUserPool(ctx, "example", &cognito.UserPoolArgs{
			AutoVerifiedAttributes: pulumi.StringArray{
				pulumi.String("email"),
			},
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewIdentityProvider(ctx, "exampleProvider", &cognito.IdentityProviderArgs{
			AttributeMapping: pulumi.StringMap{
				"email":    pulumi.String("email"),
				"username": pulumi.String("sub"),
			},
			ProviderDetails: pulumi.StringMap{
				"authorize_scopes": pulumi.String("email"),
				"client_id":        pulumi.String("your client_id"),
				"client_secret":    pulumi.String("your client_secret"),
			},
			ProviderName: pulumi.String("Google"),
			ProviderType: pulumi.String("Google"),
			UserPoolId:   example.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

