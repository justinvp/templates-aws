using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new certificate
        var test = new Aws.Dms.Certificate("test", new Aws.Dms.CertificateArgs
        {
            CertificateId = "test-dms-certificate-tf",
            CertificatePem = "...",
        });
    }

}

