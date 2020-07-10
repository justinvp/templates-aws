using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ses.DomainIdentity("example", new Aws.Ses.DomainIdentityArgs
        {
            Domain = "example.com",
        });
        var exampleAmazonsesVerificationRecord = new Aws.Route53.Record("exampleAmazonsesVerificationRecord", new Aws.Route53.RecordArgs
        {
            Name = example.Id.Apply(id => $"_amazonses.{id}"),
            Records = 
            {
                example.VerificationToken,
            },
            Ttl = 600,
            Type = "TXT",
            ZoneId = aws_route53_zone.Example.Zone_id,
        });
        var exampleVerification = new Aws.Ses.DomainIdentityVerification("exampleVerification", new Aws.Ses.DomainIdentityVerificationArgs
        {
            Domain = example.Id,
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_route53_record.example_amazonses_verification_record",
            },
        });
    }

}

