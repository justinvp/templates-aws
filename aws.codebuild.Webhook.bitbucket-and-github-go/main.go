package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codebuild"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codebuild.NewWebhook(ctx, "example", &codebuild.WebhookArgs{
			FilterGroups: codebuild.WebhookFilterGroupArray{
				&codebuild.WebhookFilterGroupArgs{
					Filter: pulumi.StringMapArray{
						pulumi.StringMap{
							"pattern": pulumi.String("PUSH"),
							"type":    pulumi.String("EVENT"),
						},
						pulumi.StringMap{
							"pattern": pulumi.String("master"),
							"type":    pulumi.String("HEAD_REF"),
						},
					},
				},
			},
			ProjectName: pulumi.Any(aws_codebuild_project.Example.Name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

