import pulumi
import pulumi_aws as aws

test_role = aws.iam.Role("testRole",
    assume_role_policy="""{
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

""",
    tags={
        "tag-key": "tag-value",
    })

