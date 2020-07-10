import pulumi
import pulumi_aws as aws

role = aws.iam.Role("role",
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
    path="/")
test_profile = aws.iam.InstanceProfile("testProfile", role=role.name)

