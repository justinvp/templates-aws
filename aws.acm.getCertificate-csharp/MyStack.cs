using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Acm.GetCertificate.InvokeAsync(new Aws.Acm.GetCertificateArgs
        {
            Domain = "tf.example.com",
            KeyTypes = 
            {
                "RSA_4096",
            },
        }));
    }

}

