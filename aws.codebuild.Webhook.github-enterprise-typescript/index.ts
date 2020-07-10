import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as github from "@pulumi/github";

const exampleWebhook = new aws.codebuild.Webhook("example", {
    projectName: aws_codebuild_project_example.name,
});
const exampleRepositoryWebhook = new github.RepositoryWebhook("example", {
    active: true,
    configuration: {
        contentType: "json",
        insecureSsl: false,
        secret: exampleWebhook.secret,
        url: exampleWebhook.payloadUrl,
    },
    events: ["push"],
    repository: github_repository_example.name,
});

