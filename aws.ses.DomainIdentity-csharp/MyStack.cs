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
            Name = "_amazonses.example.com",
            Records = 
            {
                example.VerificationToken,
            },
            Ttl = 600,
            Type = "TXT",
            ZoneId = "ABCDEFGHIJ123",
        });
    }

}

