import pulumi
import pulumi_aws as aws

test = aws.get_prefix_list(filters=[{
    "name": "prefix-list-id",
    "values": ["pl-68a54001"],
}])

