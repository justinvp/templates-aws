import pulumi
import pulumi_aws as aws

sys = aws.route53.ResolverRule("sys",
    domain_name="subdomain.example.com",
    rule_type="SYSTEM")

