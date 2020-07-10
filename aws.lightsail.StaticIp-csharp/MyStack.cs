using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.LightSail.StaticIp("test", new Aws.LightSail.StaticIpArgs
        {
        });
    }

}

