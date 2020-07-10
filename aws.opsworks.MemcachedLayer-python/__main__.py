import pulumi
import pulumi_aws as aws

cache = aws.opsworks.MemcachedLayer("cache", stack_id=aws_opsworks_stack["main"]["id"])

