package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codepipeline"
	"github.com/pulumi/pulumi-github/sdk/go/github"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		barPipeline, err := codepipeline.NewPipeline(ctx, "barPipeline", &codepipeline.PipelineArgs{
			ArtifactStore: &codepipeline.PipelineArtifactStoreArgs{
				EncryptionKey: &codepipeline.PipelineArtifactStoreEncryptionKeyArgs{
					Id:   pulumi.Any(data.Aws_kms_alias.S3kmskey.Arn),
					Type: pulumi.String("KMS"),
				},
				Location: pulumi.Any(aws_s3_bucket.Bar.Bucket),
				Type:     pulumi.String("S3"),
			},
			RoleArn: pulumi.Any(aws_iam_role.Bar.Arn),
			Stages: codepipeline.PipelineStageArray{
				&codepipeline.PipelineStageArgs{
					Action: pulumi.MapArray{
						pulumi.Map{
							"category": pulumi.String("Source"),
							"configuration": pulumi.StringMap{
								"Branch": pulumi.String("master"),
								"Owner":  pulumi.String("my-organization"),
								"Repo":   pulumi.String("test"),
							},
							"name": pulumi.String("Source"),
							"outputArtifacts": pulumi.StringArray{
								pulumi.String("test"),
							},
							"owner":    pulumi.String("ThirdParty"),
							"provider": pulumi.String("GitHub"),
							"version":  pulumi.String("1"),
						},
					},
					Name: pulumi.String("Source"),
				},
				&codepipeline.PipelineStageArgs{
					Action: pulumi.MapArray{
						pulumi.Map{
							"category": pulumi.String("Build"),
							"configuration": pulumi.StringMap{
								"ProjectName": pulumi.String("test"),
							},
							"inputArtifacts": pulumi.StringArray{
								pulumi.String("test"),
							},
							"name":     pulumi.String("Build"),
							"owner":    pulumi.String("AWS"),
							"provider": pulumi.String("CodeBuild"),
							"version":  pulumi.String("1"),
						},
					},
					Name: pulumi.String("Build"),
				},
			},
		})
		if err != nil {
			return err
		}
		webhookSecret := "super-secret"
		barWebhook, err := codepipeline.NewWebhook(ctx, "barWebhook", &codepipeline.WebhookArgs{
			Authentication: pulumi.String("GITHUB_HMAC"),
			AuthenticationConfiguration: &codepipeline.WebhookAuthenticationConfigurationArgs{
				SecretToken: pulumi.String(webhookSecret),
			},
			Filters: codepipeline.WebhookFilterArray{
				&codepipeline.WebhookFilterArgs{
					JsonPath:    pulumi.String(fmt.Sprintf("%v%v", "$", ".ref")),
					MatchEquals: pulumi.String("refs/heads/{Branch}"),
				},
			},
			TargetAction:   pulumi.String("Source"),
			TargetPipeline: barPipeline.Name,
		})
		if err != nil {
			return err
		}
		_, err = github.NewRepositoryWebhook(ctx, "barRepositoryWebhook", &github.RepositoryWebhookArgs{
			Configuration: &github.RepositoryWebhookConfigurationArgs{
				ContentType: pulumi.String("json"),
				InsecureSsl: pulumi.Bool(true),
				Secret:      pulumi.String(webhookSecret),
				Url:         barWebhook.Url,
			},
			Events: pulumi.StringArray{
				pulumi.String("push"),
			},
			Repository: pulumi.Any(github_repository.Repo.Name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

