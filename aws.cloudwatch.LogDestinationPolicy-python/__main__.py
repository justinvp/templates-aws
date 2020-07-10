import pulumi
import pulumi_aws as aws

test_destination = aws.cloudwatch.LogDestination("testDestination",
    role_arn=aws_iam_role["iam_for_cloudwatch"]["arn"],
    target_arn=aws_kinesis_stream["kinesis_for_cloudwatch"]["arn"])
test_destination_policy_policy_document = test_destination.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[{
    "actions": ["logs:PutSubscriptionFilter"],
    "effect": "Allow",
    "principals": [{
        "identifiers": ["123456789012"],
        "type": "AWS",
    }],
    "resources": [arn],
}]))
test_destination_policy_log_destination_policy = aws.cloudwatch.LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy",
    access_policy=test_destination_policy_policy_document.json,
    destination_name=test_destination.name)

