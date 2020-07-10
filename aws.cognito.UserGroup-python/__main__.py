import pulumi
import pulumi_aws as aws

main_user_pool = aws.cognito.UserPool("mainUserPool")
group_role = aws.iam.Role("groupRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Federated": "cognito-identity.amazonaws.com"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "cognito-identity.amazonaws.com:aud": "us-east-1:12345678-dead-beef-cafe-123456790ab"
        },
        "ForAnyValue:StringLike": {
          "cognito-identity.amazonaws.com:amr": "authenticated"
        }
      }
    }
  ]
}

""")
main_user_group = aws.cognito.UserGroup("mainUserGroup",
    description="Managed by Pulumi",
    precedence=42,
    role_arn=group_role.arn,
    user_pool_id=main_user_pool.id)

