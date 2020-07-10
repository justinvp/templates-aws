import pulumi
import pulumi_aws as aws

example = aws.ses.DomainIdentity("example", domain="example.com")
example_amazonses_verification_record = aws.route53.Record("exampleAmazonsesVerificationRecord",
    name="_amazonses.example.com",
    records=[example.verification_token],
    ttl="600",
    type="TXT",
    zone_id="ABCDEFGHIJ123")

