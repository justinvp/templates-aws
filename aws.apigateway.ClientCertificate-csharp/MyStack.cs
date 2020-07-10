using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var demo = new Aws.ApiGateway.ClientCertificate("demo", new Aws.ApiGateway.ClientCertificateArgs
        {
            Description = "My client certificate",
        });
    }

}

