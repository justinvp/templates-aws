import pulumi
import pulumi_aws as aws

# Example SES Domain Identity
example_domain_identity = aws.ses.DomainIdentity("exampleDomainIdentity", domain="example.com")
example_mail_from = aws.ses.MailFrom("exampleMailFrom",
    domain=example_domain_identity.domain,
    mail_from_domain=example_domain_identity.domain.apply(lambda domain: f"bounce.{domain}"))
# Example Route53 MX record
example_ses_domain_mail_from_mx = aws.route53.Record("exampleSesDomainMailFromMx",
    name=example_mail_from.mail_from_domain,
    records=["10 feedback-smtp.us-east-1.amazonses.com"],
    ttl="600",
    type="MX",
    zone_id=aws_route53_zone["example"]["id"])
# Example Route53 TXT record for SPF
example_ses_domain_mail_from_txt = aws.route53.Record("exampleSesDomainMailFromTxt",
    name=example_mail_from.mail_from_domain,
    records=["v=spf1 include:amazonses.com -all"],
    ttl="600",
    type="TXT",
    zone_id=aws_route53_zone["example"]["id"])

