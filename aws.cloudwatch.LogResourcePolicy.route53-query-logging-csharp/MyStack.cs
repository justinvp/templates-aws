using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var route53_query_logging_policyPolicyDocument = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "logs:CreateLogStream",
                        "logs:PutLogEvents",
                    },
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "route53.amazonaws.com",
                            },
                            Type = "Service",
                        },
                    },
                    Resources = 
                    {
                        "arn:aws:logs:*:*:log-group:/aws/route53/*",
                    },
                },
            },
        }));
        var route53_query_logging_policyLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy", new Aws.CloudWatch.LogResourcePolicyArgs
        {
            PolicyDocument = route53_query_logging_policyPolicyDocument.Apply(route53_query_logging_policyPolicyDocument => route53_query_logging_policyPolicyDocument.Json),
            PolicyName = "route53-query-logging-policy",
        });
    }

}

