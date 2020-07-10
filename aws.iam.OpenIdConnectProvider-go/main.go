package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewOpenIdConnectProvider(ctx, "_default", &iam.OpenIdConnectProviderArgs{
			ClientIdLists: pulumi.StringArray{
				pulumi.String("266362248691-342342xasdasdasda-apps.googleusercontent.com"),
			},
			ThumbprintLists: []interface{}{},
			Url:             pulumi.String("https://accounts.google.com"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

