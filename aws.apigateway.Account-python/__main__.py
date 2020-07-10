import pulumi
import pulumi_aws as aws

cloudwatch_role = aws.iam.Role("cloudwatchRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

""")
demo = aws.apigateway.Account("demo", cloudwatch_role_arn=cloudwatch_role.arn)
cloudwatch_role_policy = aws.iam.RolePolicy("cloudwatchRolePolicy",
    policy="""{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:DescribeLogGroups",
                "logs:DescribeLogStreams",
                "logs:PutLogEvents",
                "logs:GetLogEvents",
                "logs:FilterLogEvents"
            ],
            "Resource": "*"
        }
    ]
}

""",
    role=cloudwatch_role.id)

