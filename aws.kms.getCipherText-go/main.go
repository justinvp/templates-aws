package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		oauthConfig, err := kms.NewKey(ctx, "oauthConfig", &kms.KeyArgs{
			Description: pulumi.String("oauth config"),
			IsEnabled:   pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

