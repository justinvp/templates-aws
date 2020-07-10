import pulumi
import pulumi_aws as aws

elasticsearch_log_publishing_policy_policy_document = aws.iam.get_policy_document(statements=[{
    "actions": [
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:PutLogEventsBatch",
    ],
    "principals": [{
        "identifiers": ["es.amazonaws.com"],
        "type": "Service",
    }],
    "resources": ["arn:aws:logs:*"],
}])
elasticsearch_log_publishing_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("elasticsearch-log-publishing-policyLogResourcePolicy",
    policy_document=elasticsearch_log_publishing_policy_policy_document.json,
    policy_name="elasticsearch-log-publishing-policy")

