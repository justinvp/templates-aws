import pulumi
import pulumi_aws as aws

example = aws.organizations.Policy("example", content="""{
  "Version": "2012-10-17",
  "Statement": {
    "Effect": "Allow",
    "Action": "*",
    "Resource": "*"
  }
}

""")

