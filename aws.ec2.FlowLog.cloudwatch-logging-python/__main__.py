import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_role = aws.iam.Role("exampleRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "vpc-flow-logs.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

""")
example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
    iam_role_arn=example_role.arn,
    log_destination=example_log_group.arn,
    traffic_type="ALL",
    vpc_id=aws_vpc["example"]["id"])
example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
    policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogGroups",
        "logs:DescribeLogStreams"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}

""",
    role=example_role.id)

