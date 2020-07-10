import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket", acl="private")
firehose_role = aws.iam.Role("firehoseRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "firehose.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}

""")
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="s3",
    s3_configuration={
        "bucketArn": bucket.arn,
        "role_arn": firehose_role.arn,
    })

