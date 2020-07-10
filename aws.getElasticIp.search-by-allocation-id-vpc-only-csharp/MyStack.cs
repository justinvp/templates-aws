using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var byAllocationId = Output.Create(Aws.GetElasticIp.InvokeAsync(new Aws.GetElasticIpArgs
        {
            Id = "eipalloc-12345678",
        }));
    }

}

