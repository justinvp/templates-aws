import pulumi
import pulumi_aws as aws

example = aws.cloudtrail.Trail("example", event_selectors=[{
    "dataResource": [{
        "type": "AWS::S3::Object",
        "values": ["arn:aws:s3:::"],
    }],
    "includeManagementEvents": True,
    "readWriteType": "All",
}])

