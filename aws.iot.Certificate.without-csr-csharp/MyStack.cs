using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var cert = new Aws.Iot.Certificate("cert", new Aws.Iot.CertificateArgs
        {
            Active = true,
        });
    }

}

