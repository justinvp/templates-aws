using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new Aws.CloudWatch.LogGroupArgs
        {
            RetentionInDays = 14,
        });
        var ad_log_policyPolicyDocument = exampleLogGroup.Arn.Apply(arn => Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
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
                    Effect = "Allow",
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "ds.amazonaws.com",
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
        var ad_log_policyLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("ad-log-policyLogResourcePolicy", new Aws.CloudWatch.LogResourcePolicyArgs
        {
            PolicyDocument = ad_log_policyPolicyDocument.Apply(ad_log_policyPolicyDocument => ad_log_policyPolicyDocument.Json),
            PolicyName = "ad-log-policy",
        });
        var exampleLogService = new Aws.DirectoryService.LogService("exampleLogService", new Aws.DirectoryService.LogServiceArgs
        {
            DirectoryId = aws_directory_service_directory.Example.Id,
            LogGroupName = exampleLogGroup.Name,
        });
    }

}

