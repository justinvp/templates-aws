import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Example SES Domain Identity
const exampleDomainIdentity = new aws.ses.DomainIdentity("example", {
    domain: "example.com",
});
const exampleMailFrom = new aws.ses.MailFrom("example", {
    domain: exampleDomainIdentity.domain,
    mailFromDomain: pulumi.interpolate`bounce.${exampleDomainIdentity.domain}`,
});
// Example Route53 MX record
const exampleSesDomainMailFromMx = new aws.route53.Record("example_ses_domain_mail_from_mx", {
    name: exampleMailFrom.mailFromDomain,
    records: ["10 feedback-smtp.us-east-1.amazonses.com"], // Change to the region in which `aws_ses_domain_identity.example` is created
    ttl: 600,
    type: "MX",
    zoneId: aws_route53_zone_example.id,
});
// Example Route53 TXT record for SPF
const exampleSesDomainMailFromTxt = new aws.route53.Record("example_ses_domain_mail_from_txt", {
    name: exampleMailFrom.mailFromDomain,
    records: ["v=spf1 include:amazonses.com -all"],
    ttl: 600,
    type: "TXT",
    zoneId: aws_route53_zone_example.id,
});

