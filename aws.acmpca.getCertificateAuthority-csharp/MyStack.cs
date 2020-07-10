using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Acmpca.GetCertificateAuthority.InvokeAsync(new Aws.Acmpca.GetCertificateAuthorityArgs
        {
            Arn = "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012",
        }));
    }

}

