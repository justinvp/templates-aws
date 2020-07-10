import pulumi
import pulumi_aws as aws

user = aws.iam.User("user")
role = aws.iam.Role("role", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
group = aws.iam.Group("group")
policy = aws.iam.Policy("policy",
    description="A test policy",
    policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}

""")
test_attach = aws.iam.PolicyAttachment("test-attach",
    groups=[group.name],
    policy_arn=policy.arn,
    roles=[role.name],
    users=[user.name])

