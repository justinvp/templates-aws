import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
foo = aws.s3.Bucket("foo",
    force_destroy=True,
    policy=f"""{{
    "Version": "2012-10-17",
    "Statement": [
        {{
            "Sid": "AWSCloudTrailAclCheck",
            "Effect": "Allow",
            "Principal": {{
              "Service": "cloudtrail.amazonaws.com"
            }},
            "Action": "s3:GetBucketAcl",
            "Resource": "arn:aws:s3:::tf-test-trail"
        }},
        {{
            "Sid": "AWSCloudTrailWrite",
            "Effect": "Allow",
            "Principal": {{
              "Service": "cloudtrail.amazonaws.com"
            }},
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::tf-test-trail/prefix/AWSLogs/{current.account_id}/*",
            "Condition": {{
                "StringEquals": {{
                    "s3:x-amz-acl": "bucket-owner-full-control"
                }}
            }}
        }}
    ]
}}

""")
foobar = aws.cloudtrail.Trail("foobar",
    include_global_service_events=False,
    s3_bucket_name=foo.id,
    s3_key_prefix="prefix")

