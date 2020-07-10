import pulumi
import pulumi_aws as aws

master = aws.guardduty.Detector("master", enable=True)
bucket = aws.s3.Bucket("bucket", acl="private")
my_threat_intel_set_bucket_object = aws.s3.BucketObject("myThreatIntelSetBucketObject",
    acl="public-read",
    bucket=bucket.id,
    content="""10.0.0.0/8

""",
    key="MyThreatIntelSet")
my_threat_intel_set_threat_intel_set = aws.guardduty.ThreatIntelSet("myThreatIntelSetThreatIntelSet",
    activate=True,
    detector_id=master.id,
    format="TXT",
    location=pulumi.Output.all(my_threat_intel_set_bucket_object.bucket, my_threat_intel_set_bucket_object.key).apply(lambda bucket, key: f"https://s3.amazonaws.com/{bucket}/{key}"))

