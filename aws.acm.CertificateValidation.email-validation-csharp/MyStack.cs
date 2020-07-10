using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var certCertificate = new Aws.Acm.Certificate("certCertificate", new Aws.Acm.CertificateArgs
        {
            DomainName = "example.com",
            ValidationMethod = "EMAIL",
        });
        var certCertificateValidation = new Aws.Acm.CertificateValidation("certCertificateValidation", new Aws.Acm.CertificateValidationArgs
        {
            CertificateArn = certCertificate.Arn,
        });
    }

}

