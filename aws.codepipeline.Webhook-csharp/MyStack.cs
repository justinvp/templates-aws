using Pulumi;
using Aws = Pulumi.Aws;
using Github = Pulumi.Github;

class MyStack : Stack
{
    public MyStack()
    {
        var barPipeline = new Aws.CodePipeline.Pipeline("barPipeline", new Aws.CodePipeline.PipelineArgs
        {
            ArtifactStore = new Aws.CodePipeline.Inputs.PipelineArtifactStoreArgs
            {
                EncryptionKey = new Aws.CodePipeline.Inputs.PipelineArtifactStoreEncryptionKeyArgs
                {
                    Id = data.Aws_kms_alias.S3kmskey.Arn,
                    Type = "KMS",
                },
                Location = aws_s3_bucket.Bar.Bucket,
                Type = "S3",
            },
            RoleArn = aws_iam_role.Bar.Arn,
            Stages = 
            {
                new Aws.CodePipeline.Inputs.PipelineStageArgs
                {
                    Action = 
                    {
                        
                        {
                            { "category", "Source" },
                            { "configuration", 
                            {
                                { "Branch", "master" },
                                { "Owner", "my-organization" },
                                { "Repo", "test" },
                            } },
                            { "name", "Source" },
                            { "outputArtifacts", 
                            {
                                "test",
                            } },
                            { "owner", "ThirdParty" },
                            { "provider", "GitHub" },
                            { "version", "1" },
                        },
                    },
                    Name = "Source",
                },
                new Aws.CodePipeline.Inputs.PipelineStageArgs
                {
                    Action = 
                    {
                        
                        {
                            { "category", "Build" },
                            { "configuration", 
                            {
                                { "ProjectName", "test" },
                            } },
                            { "inputArtifacts", 
                            {
                                "test",
                            } },
                            { "name", "Build" },
                            { "owner", "AWS" },
                            { "provider", "CodeBuild" },
                            { "version", "1" },
                        },
                    },
                    Name = "Build",
                },
            },
        });
        var webhookSecret = "super-secret";
        var barWebhook = new Aws.CodePipeline.Webhook("barWebhook", new Aws.CodePipeline.WebhookArgs
        {
            Authentication = "GITHUB_HMAC",
            AuthenticationConfiguration = new Aws.CodePipeline.Inputs.WebhookAuthenticationConfigurationArgs
            {
                SecretToken = webhookSecret,
            },
            Filters = 
            {
                new Aws.CodePipeline.Inputs.WebhookFilterArgs
                {
                    JsonPath = "$.ref",
                    MatchEquals = "refs/heads/{Branch}",
                },
            },
            TargetAction = "Source",
            TargetPipeline = barPipeline.Name,
        });
        // Wire the CodePipeline webhook into a GitHub repository.
        var barRepositoryWebhook = new Github.RepositoryWebhook("barRepositoryWebhook", new Github.RepositoryWebhookArgs
        {
            Configuration = new Github.Inputs.RepositoryWebhookConfigurationArgs
            {
                ContentType = "json",
                InsecureSsl = true,
                Secret = webhookSecret,
                Url = barWebhook.Url,
            },
            Events = 
            {
                "push",
            },
            Repository = github_repository.Repo.Name,
        });
    }

}

