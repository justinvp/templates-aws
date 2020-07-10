using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var byTags = Output.Create(Aws.GetElasticIp.InvokeAsync(new Aws.GetElasticIpArgs
        {
            Tags = 
            {
                { "Name", "exampleNameTagValue" },
            },
        }));
    }

}

