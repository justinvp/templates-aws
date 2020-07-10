import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
test_user_pool = aws.cognito.UserPool("testUserPool")
test_app = aws.pinpoint.App("testApp")
test_role = aws.iam.Role("testRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "cognito-idp.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
test_role_policy = aws.iam.RolePolicy("testRolePolicy",
    policy=test_app.application_id.apply(lambda application_id: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Action": [
        "mobiletargeting:UpdateEndpoint",
        "mobiletargeting:PutItems"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:mobiletargeting:*:{current.account_id}:apps/{application_id}*"
    }}
  ]
}}

"""),
    role=test_role.id)
test_user_pool_client = aws.cognito.UserPoolClient("testUserPoolClient",
    analytics_configuration={
        "application_id": test_app.application_id,
        "external_id": "some_id",
        "role_arn": test_role.arn,
        "userDataShared": True,
    },
    user_pool_id=test_user_pool.id)

