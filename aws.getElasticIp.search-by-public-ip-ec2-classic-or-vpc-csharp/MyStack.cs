using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var byPublicIp = Output.Create(Aws.GetElasticIp.InvokeAsync(new Aws.GetElasticIpArgs
        {
            PublicIp = "1.2.3.4",
        }));
    }

}

