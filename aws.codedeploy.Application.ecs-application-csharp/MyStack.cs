using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.CodeDeploy.Application("example", new Aws.CodeDeploy.ApplicationArgs
        {
            ComputePlatform = "ECS",
        });
    }

}

