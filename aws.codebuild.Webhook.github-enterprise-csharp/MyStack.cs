using Pulumi;
using Aws = Pulumi.Aws;
using Github = Pulumi.Github;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleWebhook = new Aws.CodeBuild.Webhook("exampleWebhook", new Aws.CodeBuild.WebhookArgs
        {
            ProjectName = aws_codebuild_project.Example.Name,
        });
        var exampleRepositoryWebhook = new Github.RepositoryWebhook("exampleRepositoryWebhook", new Github.RepositoryWebhookArgs
        {
            Active = true,
            Configuration = new Github.Inputs.RepositoryWebhookConfigurationArgs
            {
                ContentType = "json",
                InsecureSsl = false,
                Secret = exampleWebhook.Secret,
                Url = exampleWebhook.PayloadUrl,
            },
            Events = 
            {
                "push",
            },
            Repository = github_repository.Example.Name,
        });
    }

}

