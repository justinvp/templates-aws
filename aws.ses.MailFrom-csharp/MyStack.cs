using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Example SES Domain Identity
        var exampleDomainIdentity = new Aws.Ses.DomainIdentity("exampleDomainIdentity", new Aws.Ses.DomainIdentityArgs
        {
            Domain = "example.com",
        });
        var exampleMailFrom = new Aws.Ses.MailFrom("exampleMailFrom", new Aws.Ses.MailFromArgs
        {
            Domain = exampleDomainIdentity.Domain,
            MailFromDomain = exampleDomainIdentity.Domain.Apply(domain => $"bounce.{domain}"),
        });
        // Example Route53 MX record
        var exampleSesDomainMailFromMx = new Aws.Route53.Record("exampleSesDomainMailFromMx", new Aws.Route53.RecordArgs
        {
            Name = exampleMailFrom.MailFromDomain,
            Records = 
            {
                "10 feedback-smtp.us-east-1.amazonses.com",
            },
            Ttl = 600,
            Type = "MX",
            ZoneId = aws_route53_zone.Example.Id,
        });
        // Example Route53 TXT record for SPF
        var exampleSesDomainMailFromTxt = new Aws.Route53.Record("exampleSesDomainMailFromTxt", new Aws.Route53.RecordArgs
        {
            Name = exampleMailFrom.MailFromDomain,
            Records = 
            {
                "v=spf1 include:amazonses.com -all",
            },
            Ttl = 600,
            Type = "TXT",
            ZoneId = aws_route53_zone.Example.Id,
        });
    }

}

