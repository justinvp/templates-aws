import pulumi
import pulumi_aws as aws

lb_user = aws.iam.User("lbUser", path="/system/")
lb_ro = aws.iam.UserPolicy("lbRo",
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

""",
    user=lb_user.name)
lb_access_key = aws.iam.AccessKey("lbAccessKey", user=lb_user.name)

