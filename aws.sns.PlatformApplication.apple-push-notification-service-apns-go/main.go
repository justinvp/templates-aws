package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sns.NewPlatformApplication(ctx, "apnsApplication", &sns.PlatformApplicationArgs{
			Platform:           pulumi.String("APNS"),
			PlatformCredential: pulumi.String("<APNS PRIVATE KEY>"),
			PlatformPrincipal:  pulumi.String("<APNS CERTIFICATE>"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

