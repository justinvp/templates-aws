import pulumi
import pulumi_aws as aws

fwd = aws.route53.ResolverRule("fwd",
    domain_name="example.com",
    resolver_endpoint_id=aws_route53_resolver_endpoint["foo"]["id"],
    rule_type="FORWARD",
    tags={
        "Environment": "Prod",
    },
    target_ips=[{
        "ip": "123.45.67.89",
    }])

