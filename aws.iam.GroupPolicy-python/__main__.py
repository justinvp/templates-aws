import pulumi
import pulumi_aws as aws

my_developers = aws.iam.Group("myDevelopers", path="/users/")
my_developer_policy = aws.iam.GroupPolicy("myDeveloperPolicy",
    group=my_developers.id,
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

