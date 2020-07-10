import pulumi
import pulumi_aws as aws

example = aws.ses.DomainIdentity("example", domain="example.com")
example_amazonses_verification_record = aws.route53.Record("exampleAmazonsesVerificationRecord",
    name=example.id.apply(lambda id: f"_amazonses.{id}"),
    records=[example.verification_token],
    ttl="600",
    type="TXT",
    zone_id=aws_route53_zone["example"]["zone_id"])
example_verification = aws.ses.DomainIdentityVerification("exampleVerification", domain=example.id,
opts=ResourceOptions(depends_on=["aws_route53_record.example_amazonses_verification_record"]))

