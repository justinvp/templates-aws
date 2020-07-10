import pulumi
import pulumi_aws as aws

main = aws.cloudtrail.get_service_account()
bucket = aws.s3.Bucket("bucket",
    force_destroy=True,
    policy=f"""{{
  "Version": "2008-10-17",
  "Statement": [
    {{
      "Sid": "Put bucket policy needed for trails",
      "Effect": "Allow",
      "Principal": {{
        "AWS": "{main.arn}"
      }},
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::tf-cloudtrail-logging-test-bucket/*"
    }},
    {{
      "Sid": "Get bucket policy needed for trails",
      "Effect": "Allow",
      "Principal": {{
        "AWS": "{main.arn}"
      }},
      "Action": "s3:GetBucketAcl",
      "Resource": "arn:aws:s3:::tf-cloudtrail-logging-test-bucket"
    }}
  ]
}}

""")

