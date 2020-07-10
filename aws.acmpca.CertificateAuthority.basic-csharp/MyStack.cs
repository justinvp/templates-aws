using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Acmpca.CertificateAuthority("example", new Aws.Acmpca.CertificateAuthorityArgs
        {
            CertificateAuthorityConfiguration = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationArgs
            {
                KeyAlgorithm = "RSA_4096",
                SigningAlgorithm = "SHA512WITHRSA",
                Subject = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs
                {
                    CommonName = "example.com",
                },
            },
            PermanentDeletionTimeInDays = 7,
        });
    }

}

