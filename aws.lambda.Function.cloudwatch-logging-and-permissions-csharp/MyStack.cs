using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testLambda = new Aws.Lambda.Function("testLambda", new Aws.Lambda.FunctionArgs
        {
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_cloudwatch_log_group.example",
                "aws_iam_role_policy_attachment.lambda_logs",
            },
        });
        // This is to optionally manage the CloudWatch Log Group for the Lambda Function.
        // If skipping this resource configuration, also add "logs:CreateLogGroup" to the IAM policy below.
        var example = new Aws.CloudWatch.LogGroup("example", new Aws.CloudWatch.LogGroupArgs
        {
            RetentionInDays = 14,
        });
        // See also the following AWS managed policy: AWSLambdaBasicExecutionRole
        var lambdaLogging = new Aws.Iam.Policy("lambdaLogging", new Aws.Iam.PolicyArgs
        {
            Description = "IAM policy for logging from a lambda",
            Path = "/",
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": [
        ""logs:CreateLogGroup"",
        ""logs:CreateLogStream"",
        ""logs:PutLogEvents""
      ],
      ""Resource"": ""arn:aws:logs:*:*:*"",
      ""Effect"": ""Allow""
    }
  ]
}

",
        });
        var lambdaLogs = new Aws.Iam.RolePolicyAttachment("lambdaLogs", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = lambdaLogging.Arn,
            Role = aws_iam_role.Iam_for_lambda.Name,
        });
    }

}

