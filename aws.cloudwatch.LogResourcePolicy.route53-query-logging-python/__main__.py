import pulumi
import pulumi_aws as aws

route53_query_logging_policy_policy_document = aws.iam.get_policy_document(statements=[{
    "actions": [
        "logs:CreateLogStream",
        "logs:PutLogEvents",
    ],
    "principals": [{
        "identifiers": ["route53.amazonaws.com"],
        "type": "Service",
    }],
    "resources": ["arn:aws:logs:*:*:log-group:/aws/route53/*"],
}])
route53_query_logging_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy",
    policy_document=route53_query_logging_policy_policy_document.json,
    policy_name="route53-query-logging-policy")

