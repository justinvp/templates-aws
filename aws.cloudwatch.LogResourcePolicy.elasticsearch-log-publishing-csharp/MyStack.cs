using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var elasticsearch_log_publishing_policyPolicyDocument = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "logs:CreateLogStream",
                        "logs:PutLogEvents",
                        "logs:PutLogEventsBatch",
                    },
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "es.amazonaws.com",
                            },
                            Type = "Service",
                        },
                    },
                    Resources = 
                    {
                        "arn:aws:logs:*",
                    },
                },
            },
        }));
        var elasticsearch_log_publishing_policyLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("elasticsearch-log-publishing-policyLogResourcePolicy", new Aws.CloudWatch.LogResourcePolicyArgs
        {
            PolicyDocument = elasticsearch_log_publishing_policyPolicyDocument.Apply(elasticsearch_log_publishing_policyPolicyDocument => elasticsearch_log_publishing_policyPolicyDocument.Json),
            PolicyName = "elasticsearch-log-publishing-policy",
        });
    }

}

