import pulumi
import pulumi_aws as aws
import pulumi_github as github

example_webhook = aws.codebuild.Webhook("exampleWebhook", project_name=aws_codebuild_project["example"]["name"])
example_repository_webhook = github.RepositoryWebhook("exampleRepositoryWebhook",
    active=True,
    configuration={
        "contentType": "json",
        "insecureSsl": False,
        "secret": example_webhook.secret,
        "url": example_webhook.payload_url,
    },
    events=["push"],
    repository=github_repository["example"]["name"])

