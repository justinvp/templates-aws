import pulumi
import pulumi_aws as aws
import pulumi_github as github

bar_pipeline = aws.codepipeline.Pipeline("barPipeline",
    artifact_store={
        "encryption_key": {
            "id": data["aws_kms_alias"]["s3kmskey"]["arn"],
            "type": "KMS",
        },
        "location": aws_s3_bucket["bar"]["bucket"],
        "type": "S3",
    },
    role_arn=aws_iam_role["bar"]["arn"],
    stages=[
        {
            "action": [{
                "category": "Source",
                "configuration": {
                    "Branch": "master",
                    "Owner": "my-organization",
                    "Repo": "test",
                },
                "name": "Source",
                "outputArtifacts": ["test"],
                "owner": "ThirdParty",
                "provider": "GitHub",
                "version": "1",
            }],
            "name": "Source",
        },
        {
            "action": [{
                "category": "Build",
                "configuration": {
                    "ProjectName": "test",
                },
                "inputArtifacts": ["test"],
                "name": "Build",
                "owner": "AWS",
                "provider": "CodeBuild",
                "version": "1",
            }],
            "name": "Build",
        },
    ])
webhook_secret = "super-secret"
bar_webhook = aws.codepipeline.Webhook("barWebhook",
    authentication="GITHUB_HMAC",
    authentication_configuration={
        "secretToken": webhook_secret,
    },
    filters=[{
        "jsonPath": "$.ref",
        "matchEquals": "refs/heads/{Branch}",
    }],
    target_action="Source",
    target_pipeline=bar_pipeline.name)
# Wire the CodePipeline webhook into a GitHub repository.
bar_repository_webhook = github.RepositoryWebhook("barRepositoryWebhook",
    configuration={
        "contentType": "json",
        "insecureSsl": True,
        "secret": webhook_secret,
        "url": bar_webhook.url,
    },
    events=["push"],
    repository=github_repository["repo"]["name"])

