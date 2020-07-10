import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
identity = aws.ses.DomainIdentity("identity", domain="example.com")
role = aws.iam.Role("role", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "pinpoint.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
email = aws.pinpoint.EmailChannel("email",
    application_id=app.application_id,
    from_address="user@example.com",
    identity=identity.arn,
    role_arn=role.arn)
role_policy = aws.iam.RolePolicy("rolePolicy",
    policy="""{
  "Version": "2012-10-17",
  "Statement": {
    "Action": [
      "mobileanalytics:PutEvents",
      "mobileanalytics:PutItems"
    ],
    "Effect": "Allow",
    "Resource": [
      "*"
    ]
  }
}

""",
    role=role.id)

