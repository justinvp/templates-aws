using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var us_east_1 = new Aws.Provider("us-east-1", new Aws.ProviderArgs
        {
            Region = "us-east-1",
        });
        var awsRoute53ExampleCom = new Aws.CloudWatch.LogGroup("awsRoute53ExampleCom", new Aws.CloudWatch.LogGroupArgs
        {
            RetentionInDays = 30,
        }, new CustomResourceOptions
        {
            Provider = "aws.us-east-1",
        });
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
        }, new CustomResourceOptions
        {
            Provider = "aws.us-east-1",
        });
        var exampleComZone = new Aws.Route53.Zone("exampleComZone", new Aws.Route53.ZoneArgs
        {
        });
        var exampleComQueryLog = new Aws.Route53.QueryLog("exampleComQueryLog", new Aws.Route53.QueryLogArgs
        {
            CloudwatchLogGroupArn = awsRoute53ExampleCom.Arn,
            ZoneId = exampleComZone.ZoneId,
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_cloudwatch_log_resource_policy.route53-query-logging-policy",
            },
        });
    }

}

