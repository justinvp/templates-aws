import pulumi
import pulumi_aws as aws

stream = aws.kinesis.get_stream(name="stream-name")

