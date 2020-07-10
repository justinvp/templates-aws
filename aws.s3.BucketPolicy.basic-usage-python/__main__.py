import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket")
bucket_policy = aws.s3.BucketPolicy("bucketPolicy",
    bucket=bucket.id,
    policy="""{
  "Version": "2012-10-17",
  "Id": "MYBUCKETPOLICY",
  "Statement": [
    {
      "Sid": "IPAllow",
      "Effect": "Deny",
      "Principal": "*",
      "Action": "s3:*",
      "Resource": "arn:aws:s3:::my_tf_test_bucket/*",
      "Condition": {
         "IpAddress": {"aws:SourceIp": "8.8.8.8/32"}
      }
    }
  ]
}

""")

