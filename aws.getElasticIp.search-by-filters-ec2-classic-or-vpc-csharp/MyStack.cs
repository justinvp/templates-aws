using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var byFilter = Output.Create(Aws.GetElasticIp.InvokeAsync(new Aws.GetElasticIpArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetElasticIpFilterArgs
                {
                    Name = "tag:Name",
                    Values = 
                    {
                        "exampleNameTagValue",
                    },
                },
            },
        }));
    }

}

