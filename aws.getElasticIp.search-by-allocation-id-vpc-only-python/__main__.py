import pulumi
import pulumi_aws as aws

by_allocation_id = aws.get_elastic_ip(id="eipalloc-12345678")

