import pulumi
import pulumi_aws as aws

lb_user = aws.iam.User("lbUser", path="/system/")
lb_access_key = aws.iam.AccessKey("lbAccessKey",
    pgp_key="keybase:some_person_that_exists",
    user=lb_user.name)
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
pulumi.export("secret", lb_access_key.encrypted_secret)

