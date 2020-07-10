import pulumi
import pulumi_aws as aws

policy = aws.iam.Policy("policy",
    description="My test policy",
    path="/",
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

