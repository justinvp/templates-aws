using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.Sns.Topic("test", new Aws.Sns.TopicArgs
        {
        });
        var snsTopicPolicy = test.Arn.Apply(arn => Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            PolicyId = "__default_policy_ID",
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "SNS:Subscribe",
                        "SNS:SetTopicAttributes",
                        "SNS:RemovePermission",
                        "SNS:Receive",
                        "SNS:Publish",
                        "SNS:ListSubscriptionsByTopic",
                        "SNS:GetTopicAttributes",
                        "SNS:DeleteTopic",
                        "SNS:AddPermission",
                    },
                    Condition = 
                    {
                        
                        {
                            { "test", "StringEquals" },
                            { "values", 
                            {
                                @var.Account_id,
                            } },
                            { "variable", "AWS:SourceOwner" },
                        },
                    },
                    Effect = "Allow",
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "*",
                            },
                            Type = "AWS",
                        },
                    },
                    Resources = 
                    {
                        arn,
                    },
                    Sid = "__default_statement_ID",
                },
            },
        }));
        var @default = new Aws.Sns.TopicPolicy("default", new Aws.Sns.TopicPolicyArgs
        {
            Arn = test.Arn,
            Policy = snsTopicPolicy.Apply(snsTopicPolicy => snsTopicPolicy.Json),
        });
    }

}

