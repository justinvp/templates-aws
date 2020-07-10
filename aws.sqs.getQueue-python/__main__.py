import pulumi
import pulumi_aws as aws

example = aws.sqs.get_queue(name="queue")

