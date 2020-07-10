import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

us_east_1 = pulumi.providers.Aws("us-east-1", region="us-east-1")
aws_route53_example_com = aws.cloudwatch.LogGroup("awsRoute53ExampleCom", retention_in_days=30,
opts=ResourceOptions(provider="aws.us-east-1"))
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
    policy_name="route53-query-logging-policy",
    opts=ResourceOptions(provider="aws.us-east-1"))
example_com_zone = aws.route53.Zone("exampleComZone")
example_com_query_log = aws.route53.QueryLog("exampleComQueryLog",
    cloudwatch_log_group_arn=aws_route53_example_com.arn,
    zone_id=example_com_zone.zone_id,
    opts=ResourceOptions(depends_on=["aws_cloudwatch_log_resource_policy.route53-query-logging-policy"]))

