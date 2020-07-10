using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var console = new Aws.CloudWatch.EventRule("console", new Aws.CloudWatch.EventRuleArgs
        {
            Description = "Capture each AWS Console Sign In",
            EventPattern = @"{
  ""detail-type"": [
    ""AWS Console Sign In via CloudTrail""
  ]
}

",
        });
        var awsLogins = new Aws.Sns.Topic("awsLogins", new Aws.Sns.TopicArgs
        {
        });
        var sns = new Aws.CloudWatch.EventTarget("sns", new Aws.CloudWatch.EventTargetArgs
        {
            Arn = awsLogins.Arn,
            Rule = console.Name,
        });
        var snsTopicPolicy = awsLogins.Arn.Apply(arn => Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "SNS:Publish",
                    },
                    Effect = "Allow",
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "events.amazonaws.com",
                            },
                            Type = "Service",
                        },
                    },
                    Resources = 
                    {
                        arn,
                    },
                },
            },
        }));
        var @default = new Aws.Sns.TopicPolicy("default", new Aws.Sns.TopicPolicyArgs
        {
            Arn = awsLogins.Arn,
            Policy = snsTopicPolicy.Apply(snsTopicPolicy => snsTopicPolicy.Json),
        });
    }

}

