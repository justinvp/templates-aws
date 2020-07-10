import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// This is to optionally manage the CloudWatch Log Group for the Lambda Function.
// If skipping this resource configuration, also add "logs:CreateLogGroup" to the IAM policy below.
const example = new aws.cloudwatch.LogGroup("example", {
    retentionInDays: 14,
});
// See also the following AWS managed policy: AWSLambdaBasicExecutionRole
const lambdaLogging = new aws.iam.Policy("lambda_logging", {
    description: "IAM policy for logging from a lambda",
    path: "/",
    policy: `{
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
`,
});
const lambdaLogs = new aws.iam.RolePolicyAttachment("lambda_logs", {
    policyArn: lambdaLogging.arn,
    role: aws_iam_role_iam_for_lambda.name,
});
const testLambda = new aws.lambda.Function("test_lambda", {}, { dependsOn: [example, lambdaLogs] });

