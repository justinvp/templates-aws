import pulumi
import pulumi_aws as aws

hoge_bucket = aws.s3.Bucket("hogeBucket", region="us-east-1")
hoge_bucket_policy = aws.s3.BucketPolicy("hogeBucketPolicy",
    bucket=hoge_bucket.bucket,
    policy="""{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "SSMBucketPermissionsCheck",
            "Effect": "Allow",
            "Principal": {
                "Service": "ssm.amazonaws.com"
            },
            "Action": "s3:GetBucketAcl",
            "Resource": "arn:aws:s3:::tf-test-bucket-1234"
        },
        {
            "Sid": " SSMBucketDelivery",
            "Effect": "Allow",
            "Principal": {
                "Service": "ssm.amazonaws.com"
            },
            "Action": "s3:PutObject",
            "Resource": ["arn:aws:s3:::tf-test-bucket-1234/*"],
            "Condition": {
                "StringEquals": {
                    "s3:x-amz-acl": "bucket-owner-full-control"
                }
            }
        }
    ]
}

""")
foo = aws.ssm.ResourceDataSync("foo", s3_destination={
    "bucket_name": hoge_bucket.bucket,
    "region": hoge_bucket.region,
})

