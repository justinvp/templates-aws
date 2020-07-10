using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.CodeBuild.Webhook("example", new Aws.CodeBuild.WebhookArgs
        {
            FilterGroups = 
            {
                new Aws.CodeBuild.Inputs.WebhookFilterGroupArgs
                {
                    Filter = 
                    {
                        
                        {
                            { "pattern", "PUSH" },
                            { "type", "EVENT" },
                        },
                        
                        {
                            { "pattern", "master" },
                            { "type", "HEAD_REF" },
                        },
                    },
                },
            },
            ProjectName = aws_codebuild_project.Example.Name,
        });
    }

}

