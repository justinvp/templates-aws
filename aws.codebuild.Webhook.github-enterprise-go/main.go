package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codebuild"
	"github.com/pulumi/pulumi-github/sdk/go/github"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleWebhook, err := codebuild.NewWebhook(ctx, "exampleWebhook", &codebuild.WebhookArgs{
			ProjectName: pulumi.Any(aws_codebuild_project.Example.Name),
		})
		if err != nil {
			return err
		}
		_, err = github.NewRepositoryWebhook(ctx, "exampleRepositoryWebhook", &github.RepositoryWebhookArgs{
			Active: pulumi.Bool(true),
			Configuration: &github.RepositoryWebhookConfigurationArgs{
				ContentType: pulumi.String("json"),
				InsecureSsl: pulumi.Bool(false),
				Secret:      exampleWebhook.Secret,
				Url:         exampleWebhook.PayloadUrl,
			},
			Events: pulumi.StringArray{
				pulumi.String("push"),
			},
			Repository: pulumi.Any(github_repository.Example.Name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

