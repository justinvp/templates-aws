import pulumi
import pulumi_aws as aws

my_domain = aws.elasticsearch.get_domain(domain_name="my-domain-name")

