import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_rule(domain_name="subdomain.example.com",
    rule_type="SYSTEM")

