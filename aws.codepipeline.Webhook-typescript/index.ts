import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as github from "@pulumi/github";

const barPipeline = new aws.codepipeline.Pipeline("bar", {
    artifactStores: {
        encryptionKey: {
            id: aws_kms_alias_s3kmskey.arn,
            type: "KMS",
        },
        location: aws_s3_bucket_bar.bucket,
        type: "S3",
    },
    roleArn: aws_iam_role_bar.arn,
    stages: [
        {
            actions: [{
                category: "Source",
                configuration: {
                    Branch: "master",
                    Owner: "my-organization",
                    Repo: "test",
                },
                name: "Source",
                outputArtifacts: ["test"],
                owner: "ThirdParty",
                provider: "GitHub",
                version: "1",
            }],
            name: "Source",
        },
        {
            actions: [{
                category: "Build",
                configuration: {
                    ProjectName: "test",
                },
                inputArtifacts: ["test"],
                name: "Build",
                owner: "AWS",
                provider: "CodeBuild",
                version: "1",
            }],
            name: "Build",
        },
    ],
});
const webhookSecret = "super-secret";
const barWebhook = new aws.codepipeline.Webhook("bar", {
    authentication: "GITHUB_HMAC",
    authenticationConfiguration: {
        secretToken: webhookSecret,
    },
    filters: [{
        jsonPath: "$.ref",
        matchEquals: "refs/heads/{Branch}",
    }],
    targetAction: "Source",
    targetPipeline: barPipeline.name,
});
// Wire the CodePipeline webhook into a GitHub repository.
const barRepositoryWebhook = new github.RepositoryWebhook("bar", {
    configuration: {
        contentType: "json",
        insecureSsl: true,
        secret: webhookSecret,
        url: barWebhook.url,
    },
    events: ["push"],
    repository: github_repository_repo.name,
});

