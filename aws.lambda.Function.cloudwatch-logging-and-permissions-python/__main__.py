import pulumi
import pulumi_aws as aws

test_lambda = aws.lambda_.Function("testLambda", opts=ResourceOptions(depends_on=[
        "aws_cloudwatch_log_group.example",
        "aws_iam_role_policy_attachment.lambda_logs",
    ]))
# This is to optionally manage the CloudWatch Log Group for the Lambda Function.
# If skipping this resource configuration, also add "logs:CreateLogGroup" to the IAM policy below.
example = aws.cloudwatch.LogGroup("example", retention_in_days=14)
# See also the following AWS managed policy: AWSLambdaBasicExecutionRole
lambda_logging = aws.iam.Policy("lambdaLogging",
    description="IAM policy for logging from a lambda",
    path="/",
    policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}

""")
lambda_logs = aws.iam.RolePolicyAttachment("lambdaLogs",
    policy_arn=lambda_logging.arn,
    role=aws_iam_role["iam_for_lambda"]["name"])

