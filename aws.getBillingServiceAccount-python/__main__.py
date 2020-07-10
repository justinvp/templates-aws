import pulumi
import pulumi_aws as aws

main = aws.get_billing_service_account()
billing_logs = aws.s3.Bucket("billingLogs",
    acl="private",
    policy=f"""{{
  "Id": "Policy",
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Action": [
        "s3:GetBucketAcl", "s3:GetBucketPolicy"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::my-billing-tf-test-bucket",
      "Principal": {{
        "AWS": [
          "{main.arn}"
        ]
      }}
    }},
    {{
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::my-billing-tf-test-bucket/*",
      "Principal": {{
        "AWS": [
          "{main.arn}"
        ]
      }}
    }}
  ]
}}

""")

