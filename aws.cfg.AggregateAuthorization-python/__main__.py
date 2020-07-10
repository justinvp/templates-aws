import pulumi
import pulumi_aws as aws

example = aws.cfg.AggregateAuthorization("example",
    account_id="123456789012",
    region="eu-west-2")

