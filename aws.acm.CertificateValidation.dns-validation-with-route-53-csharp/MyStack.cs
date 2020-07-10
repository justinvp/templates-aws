using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var certCertificate = new Aws.Acm.Certificate("certCertificate", new Aws.Acm.CertificateArgs
        {
            DomainName = "example.com",
            ValidationMethod = "DNS",
        });
        var zone = Output.Create(Aws.Route53.GetZone.InvokeAsync(new Aws.Route53.GetZoneArgs
        {
            Name = "example.com.",
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
        var certCertificateValidation = new Aws.Acm.CertificateValidation("certCertificateValidation", new Aws.Acm.CertificateValidationArgs
        {
            CertificateArn = certCertificate.Arn,
            ValidationRecordFqdns = 
            {
                certValidation.Fqdn,
            },
        });
        var frontEnd = new Aws.LB.Listener("frontEnd", new Aws.LB.ListenerArgs
        {
            CertificateArn = certCertificateValidation.CertificateArn,
        });
    }

}

