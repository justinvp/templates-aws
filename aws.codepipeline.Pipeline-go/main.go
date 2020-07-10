package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codepipeline"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		codepipelineBucket, err := s3.NewBucket(ctx, "codepipelineBucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		codepipelineRole, err := iam.NewRole(ctx, "codepipelineRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"codepipeline.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "codepipelinePolicy", &iam.RolePolicyArgs{
			Policy: pulumi.All(codepipelineBucket.Arn, codepipelineBucket.Arn).ApplyT(func(_args []interface{}) (string, error) {
				codepipelineBucketArn := _args[0].(string)
				codepipelineBucketArn1 := _args[1].(string)
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\":\"Allow\",\n", "      \"Action\": [\n", "        \"s3:GetObject\",\n", "        \"s3:GetObjectVersion\",\n", "        \"s3:GetBucketVersioning\",\n", "        \"s3:PutObject\"\n", "      ],\n", "      \"Resource\": [\n", "        \"", codepipelineBucketArn, "\",\n", "        \"", codepipelineBucketArn1, "/*\"\n", "      ]\n", "    },\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Action\": [\n", "        \"codebuild:BatchGetBuilds\",\n", "        \"codebuild:StartBuild\"\n", "      ],\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
			Role: codepipelineRole.ID(),
		})
		if err != nil {
			return err
		}
		s3kmskey, err := kms.LookupAlias(ctx, &kms.LookupAliasArgs{
			Name: "alias/myKmsKey",
		}, nil)
		if err != nil {
			return err
		}
		_, err = codepipeline.NewPipeline(ctx, "codepipeline", &codepipeline.PipelineArgs{
			ArtifactStore: &codepipeline.PipelineArtifactStoreArgs{
				EncryptionKey: &codepipeline.PipelineArtifactStoreEncryptionKeyArgs{
					Id:   pulumi.String(s3kmskey.Arn),
					Type: pulumi.String("KMS"),
				},
				Location: codepipelineBucket.Bucket,
				Type:     pulumi.String("S3"),
			},
			RoleArn: codepipelineRole.Arn,
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
								pulumi.String("source_output"),
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
								pulumi.String("source_output"),
							},
							"name": pulumi.String("Build"),
							"outputArtifacts": pulumi.StringArray{
								pulumi.String("build_output"),
							},
							"owner":    pulumi.String("AWS"),
							"provider": pulumi.String("CodeBuild"),
							"version":  pulumi.String("1"),
						},
					},
					Name: pulumi.String("Build"),
				},
				&codepipeline.PipelineStageArgs{
					Action: pulumi.MapArray{
						pulumi.Map{
							"category": pulumi.String("Deploy"),
							"configuration": pulumi.StringMap{
								"ActionMode":     pulumi.String("REPLACE_ON_FAILURE"),
								"Capabilities":   pulumi.String("CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM"),
								"OutputFileName": pulumi.String("CreateStackOutput.json"),
								"StackName":      pulumi.String("MyStack"),
								"TemplatePath":   pulumi.String("build_output::sam-templated.yaml"),
							},
							"inputArtifacts": pulumi.StringArray{
								pulumi.String("build_output"),
							},
							"name":     pulumi.String("Deploy"),
							"owner":    pulumi.String("AWS"),
							"provider": pulumi.String("CloudFormation"),
							"version":  pulumi.String("1"),
						},
					},
					Name: pulumi.String("Deploy"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

