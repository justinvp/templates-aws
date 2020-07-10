import pulumi
import pulumi_aws as aws

example = aws.cloudtrail.Trail("example", event_selectors=[{
    "dataResource": [{
        "type": "AWS::Lambda::Function",
        "values": ["arn:aws:lambda"],
    }],
    "includeManagementEvents": True,
    "readWriteType": "All",
}])

