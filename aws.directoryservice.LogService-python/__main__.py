import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=14)
ad_log_policy_policy_document = example_log_group.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[{
    "actions": [
        "logs:CreateLogStream",
        "logs:PutLogEvents",
    ],
    "effect": "Allow",
    "principals": [{
        "identifiers": ["ds.amazonaws.com"],
        "type": "Service",
    }],
    "resources": [arn],
}]))
ad_log_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("ad-log-policyLogResourcePolicy",
    policy_document=ad_log_policy_policy_document.json,
    policy_name="ad-log-policy")
example_log_service = aws.directoryservice.LogService("exampleLogService",
    directory_id=aws_directory_service_directory["example"]["id"],
    log_group_name=example_log_group.name)

