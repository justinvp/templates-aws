import pulumi
import pulumi_aws as aws

example = aws.rds.Instance("example",
    allocated_storage=50,
    max_allocated_storage=100)

