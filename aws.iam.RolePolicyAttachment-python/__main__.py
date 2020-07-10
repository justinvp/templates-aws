import pulumi
import pulumi_aws as aws

role = aws.iam.Role("role", assume_role_policy="""    {
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
test_attach = aws.iam.RolePolicyAttachment("test-attach",
    policy_arn=policy.arn,
    role=role.name)

