using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testDestination = new Aws.CloudWatch.LogDestination("testDestination", new Aws.CloudWatch.LogDestinationArgs
        {
            RoleArn = aws_iam_role.Iam_for_cloudwatch.Arn,
            TargetArn = aws_kinesis_stream.Kinesis_for_cloudwatch.Arn,
        });
        var testDestinationPolicyPolicyDocument = testDestination.Arn.Apply(arn => Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "logs:PutSubscriptionFilter",
                    },
                    Effect = "Allow",
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "123456789012",
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
        }));
        var testDestinationPolicyLogDestinationPolicy = new Aws.CloudWatch.LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy", new Aws.CloudWatch.LogDestinationPolicyArgs
        {
            AccessPolicy = testDestinationPolicyPolicyDocument.Apply(testDestinationPolicyPolicyDocument => testDestinationPolicyPolicyDocument.Json),
            DestinationName = testDestination.Name,
        });
    }

}

