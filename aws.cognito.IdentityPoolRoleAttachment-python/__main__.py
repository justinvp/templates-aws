import pulumi
import pulumi_aws as aws

main_identity_pool = aws.cognito.IdentityPool("mainIdentityPool",
    allow_unauthenticated_identities=False,
    identity_pool_name="identity pool",
    supported_login_providers={
        "graph.facebook.com": "7346241598935555",
    })
authenticated_role = aws.iam.Role("authenticatedRole", assume_role_policy=main_identity_pool.id.apply(lambda id: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Effect": "Allow",
      "Principal": {{
        "Federated": "cognito-identity.amazonaws.com"
      }},
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {{
        "StringEquals": {{
          "cognito-identity.amazonaws.com:aud": "{id}"
        }},
        "ForAnyValue:StringLike": {{
          "cognito-identity.amazonaws.com:amr": "authenticated"
        }}
      }}
    }}
  ]
}}

"""))
authenticated_role_policy = aws.iam.RolePolicy("authenticatedRolePolicy",
    policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "mobileanalytics:PutEvents",
        "cognito-sync:*",
        "cognito-identity:*"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}

""",
    role=authenticated_role.id)
main_identity_pool_role_attachment = aws.cognito.IdentityPoolRoleAttachment("mainIdentityPoolRoleAttachment",
    identity_pool_id=main_identity_pool.id,
    role_mappings=[{
        "ambiguousRoleResolution": "AuthenticatedRole",
        "identity_provider": "graph.facebook.com",
        "mappingRule": [{
            "claim": "isAdmin",
            "matchType": "Equals",
            "role_arn": authenticated_role.arn,
            "value": "paid",
        }],
        "type": "Rules",
    }],
    roles={
        "authenticated": authenticated_role.arn,
    })

