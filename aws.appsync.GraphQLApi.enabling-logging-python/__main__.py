import pulumi
import pulumi_aws as aws

example_role = aws.iam.Role("exampleRole", assume_role_policy="""{
    "Version": "2012-10-17",
    "Statement": [
        {
        "Effect": "Allow",
        "Principal": {
            "Service": "appsync.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
        }
    ]
}

""")
example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs",
    role=example_role.name)
example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", log_config={
    "cloudwatchLogsRoleArn": example_role.arn,
    "fieldLogLevel": "ERROR",
})

