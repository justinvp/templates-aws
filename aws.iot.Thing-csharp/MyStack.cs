using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Iot.Thing("example", new Aws.Iot.ThingArgs
        {
            Attributes = 
            {
                { "First", "examplevalue" },
            },
        });
    }

}

