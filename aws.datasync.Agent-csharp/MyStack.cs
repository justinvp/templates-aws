using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DataSync.Agent("example", new Aws.DataSync.AgentArgs
        {
            IpAddress = "1.2.3.4",
        });
    }

}

