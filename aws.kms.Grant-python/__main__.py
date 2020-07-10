import pulumi
import pulumi_aws as aws

key = aws.kms.Key("key")
role = aws.iam.Role("role", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
grant = aws.kms.Grant("grant",
    constraints=[{
        "encryptionContextEquals": {
            "Department": "Finance",
        },
    }],
    grantee_principal=role.arn,
    key_id=key.key_id,
    operations=[
        "Encrypt",
        "Decrypt",
        "GenerateDataKey",
    ])

