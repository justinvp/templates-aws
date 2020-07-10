import pulumi
import pulumi_aws as aws

example_bucket = aws.s3.Bucket("exampleBucket")
example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
    log_destination=example_bucket.arn,
    log_destination_type="s3",
    traffic_type="ALL",
    vpc_id=aws_vpc["example"]["id"])

