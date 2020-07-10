import pulumi
import pulumi_aws as aws

hoge_bucket = aws.s3.Bucket("hogeBucket")
hoge_database = aws.athena.Database("hogeDatabase",
    bucket=hoge_bucket.bucket,
    name="database_name")

