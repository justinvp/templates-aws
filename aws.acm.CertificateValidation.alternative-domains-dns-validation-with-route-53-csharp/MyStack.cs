using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var certCertificate = new Aws.Acm.Certificate("certCertificate", new Aws.Acm.CertificateArgs
        {
            DomainName = "example.com",
            SubjectAlternativeNames = 
            {
                "www.example.com",
                "example.org",
            },
            ValidationMethod = "DNS",
        });
        var zone = Output.Create(Aws.Route53.GetZone.InvokeAsync(new Aws.Route53.GetZoneArgs
        {
            Name = "example.com.",
            PrivateZone = false,
        }));
        var zoneAlt = Output.Create(Aws.Route53.GetZone.InvokeAsync(new Aws.Route53.GetZoneArgs
        {
            Name = "example.org.",
            PrivateZone = false,
        }));
        var certValidation = new Aws.Route53.Record("certValidation", new Aws.Route53.RecordArgs
        {
            Name = certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[0].ResourceRecordName),
            Records = 
            {
                certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[0].ResourceRecordValue),
            },
            Ttl = 60,
            Type = certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[0].ResourceRecordType),
            ZoneId = zone.Apply(zone => zone.ZoneId),
        });
        var certValidationAlt1 = new Aws.Route53.Record("certValidationAlt1", new Aws.Route53.RecordArgs
        {
            Name = certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[1].ResourceRecordName),
            Records = 
            {
                certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[1].ResourceRecordValue),
            },
            Ttl = 60,
            Type = certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[1].ResourceRecordType),
            ZoneId = zone.Apply(zone => zone.ZoneId),
        });
        var certValidationAlt2 = new Aws.Route53.Record("certValidationAlt2", new Aws.Route53.RecordArgs
        {
            Name = certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[2].ResourceRecordName),
            Records = 
            {
                certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[2].ResourceRecordValue),
            },
            Ttl = 60,
            Type = certCertificate.DomainValidationOptions.Apply(domainValidationOptions => domainValidationOptions[2].ResourceRecordType),
            ZoneId = zoneAlt.Apply(zoneAlt => zoneAlt.ZoneId),
        });
        var certCertificateValidation = new Aws.Acm.CertificateValidation("certCertificateValidation", new Aws.Acm.CertificateValidationArgs
        {
            CertificateArn = certCertificate.Arn,
            ValidationRecordFqdns = 
            {
                certValidation.Fqdn,
                certValidationAlt1.Fqdn,
                certValidationAlt2.Fqdn,
            },
        });
        var frontEnd = new Aws.LB.Listener("frontEnd", new Aws.LB.ListenerArgs
        {
            CertificateArn = certCertificateValidation.CertificateArn,
        });
    }

}

