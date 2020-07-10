package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codebuild"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codebuild.NewSourceCredential(ctx, "example", &codebuild.SourceCredentialArgs{
			AuthType:   pulumi.String("PERSONAL_ACCESS_TOKEN"),
			ServerType: pulumi.String("GITHUB"),
			Token:      pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

