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
		_, err = kms.NewCiphertext(ctx, "oauth", &kms.CiphertextArgs{
			KeyId:     oauthConfig.KeyId,
			Plaintext: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\n", "  \"client_id\": \"e587dbae22222f55da22\",\n", "  \"client_secret\": \"8289575d00000ace55e1815ec13673955721b8a5\"\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

