import pulumi
import pulumi_aws as aws

master = aws.guardduty.Detector("master", enable=True)
bucket = aws.s3.Bucket("bucket", acl="private")
my_ip_set_bucket_object = aws.s3.BucketObject("myIPSetBucketObject",
    acl="public-read",
    bucket=bucket.id,
    content="""10.0.0.0/8

""",
    key="MyIPSet")
my_ip_set_ip_set = aws.guardduty.IPSet("myIPSetIPSet",
    activate=True,
    detector_id=master.id,
    format="TXT",
    location=pulumi.Output.all(my_ip_set_bucket_object.bucket, my_ip_set_bucket_object.key).apply(lambda bucket, key: f"https://s3.amazonaws.com/{bucket}/{key}"))

