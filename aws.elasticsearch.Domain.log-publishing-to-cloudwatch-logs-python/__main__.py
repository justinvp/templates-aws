import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_log_resource_policy = aws.cloudwatch.LogResourcePolicy("exampleLogResourcePolicy",
    policy_document="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "es.amazonaws.com"
      },
      "Action": [
        "logs:PutLogEvents",
        "logs:PutLogEventsBatch",
        "logs:CreateLogStream"
      ],
      "Resource": "arn:aws:logs:*"
    }
  ]
}

""",
    policy_name="example")
example_domain = aws.elasticsearch.Domain("exampleDomain", log_publishing_options=[{
    "cloudwatch_log_group_arn": example_log_group.arn,
    "logType": "INDEX_SLOW_LOGS",
}])

