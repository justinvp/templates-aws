using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var cert = new Aws.Acm.Certificate("cert", new Aws.Acm.CertificateArgs
        {
            DomainName = "example.com",
            Tags = 
            {
                { "Environment", "test" },
            },
            ValidationMethod = "DNS",
        });
    }

}

