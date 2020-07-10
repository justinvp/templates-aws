import pulumi
import pulumi_aws as aws

example_domain_name = aws.apigateway.DomainName("exampleDomainName",
    certificate_arn=aws_acm_certificate_validation["example"]["certificate_arn"],
    domain_name="api.example.com")
# Example DNS record using Route53.
# Route53 is not specifically required; any DNS host can be used.
example_record = aws.route53.Record("exampleRecord",
    aliases=[{
        "evaluateTargetHealth": True,
        "name": example_domain_name.cloudfront_domain_name,
        "zone_id": example_domain_name.cloudfront_zone_id,
    }],
    name=example_domain_name.domain_name,
    type="A",
    zone_id=aws_route53_zone["example"]["id"])

