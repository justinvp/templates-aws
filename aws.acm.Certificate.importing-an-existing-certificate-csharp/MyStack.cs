using Pulumi;
using Aws = Pulumi.Aws;
using Tls = Pulumi.Tls;

class MyStack : Stack
{
    public MyStack()
    {
        var examplePrivateKey = new Tls.PrivateKey("examplePrivateKey", new Tls.PrivateKeyArgs
        {
            Algorithm = "RSA",
        });
        var exampleSelfSignedCert = new Tls.SelfSignedCert("exampleSelfSignedCert", new Tls.SelfSignedCertArgs
        {
            AllowedUses = 
            {
                "key_encipherment",
                "digital_signature",
                "server_auth",
            },
            KeyAlgorithm = "RSA",
            PrivateKeyPem = examplePrivateKey.PrivateKeyPem,
            Subjects = 
            {
                new Tls.Inputs.SelfSignedCertSubjectArgs
                {
                    CommonName = "example.com",
                    Organization = "ACME Examples, Inc",
                },
            },
            ValidityPeriodHours = 12,
        });
        var cert = new Aws.Acm.Certificate("cert", new Aws.Acm.CertificateArgs
        {
            CertificateBody = exampleSelfSignedCert.CertPem,
            PrivateKey = examplePrivateKey.PrivateKeyPem,
        });
    }

}

