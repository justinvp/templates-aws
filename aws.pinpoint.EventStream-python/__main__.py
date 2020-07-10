import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
test_stream = aws.kinesis.Stream("testStream", shard_count=1)
test_role = aws.iam.Role("testRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "pinpoint.us-east-1.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
stream = aws.pinpoint.EventStream("stream",
    application_id=app.application_id,
    destination_stream_arn=test_stream.arn,
    role_arn=test_role.arn)
test_role_policy = aws.iam.RolePolicy("testRolePolicy",
    policy="""{
  "Version": "2012-10-17",
  "Statement": {
    "Action": [
      "kinesis:PutRecords",
      "kinesis:DescribeStream"
    ],
    "Effect": "Allow",
    "Resource": [
      "arn:aws:kinesis:us-east-1:*:*/*"
    ]
  }
}

""",
    role=test_role.id)

