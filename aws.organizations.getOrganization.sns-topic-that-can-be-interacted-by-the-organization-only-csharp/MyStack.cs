using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Organizations.GetOrganization.InvokeAsync());
        var snsTopic = new Aws.Sns.Topic("snsTopic", new Aws.Sns.TopicArgs
        {
        });
        var snsTopicPolicyPolicyDocument = Output.Tuple(example, snsTopic.Arn).Apply(values =>
        {
            var example = values.Item1;
            var arn = values.Item2;
            return Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
            {
                Statements = 
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                    {
                        Actions = 
                        {
                            "SNS:Subscribe",
                            "SNS:Publish",
                        },
                        Condition = 
                        {
                            
                            {
                                { "test", "StringEquals" },
                                { "values", 
                                {
                                    example.Id,
                                } },
                                { "variable", "aws:PrincipalOrgID" },
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
                    },
                },
            });
        });
        var snsTopicPolicyTopicPolicy = new Aws.Sns.TopicPolicy("snsTopicPolicyTopicPolicy", new Aws.Sns.TopicPolicyArgs
        {
            Arn = snsTopic.Arn,
            Policy = snsTopicPolicyPolicyDocument.Apply(snsTopicPolicyPolicyDocument => snsTopicPolicyPolicyDocument.Json),
        });
    }

}

