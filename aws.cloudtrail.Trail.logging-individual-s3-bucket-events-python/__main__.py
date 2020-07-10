import pulumi
import pulumi_aws as aws

important_bucket = aws.s3.get_bucket(bucket="important-bucket")
example = aws.cloudtrail.Trail("example", event_selectors=[{
    "dataResource": [{
        "type": "AWS::S3::Object",
        "values": [f"{important_bucket.arn}/"],
    }],
    "includeManagementEvents": True,
    "readWriteType": "All",
}])

