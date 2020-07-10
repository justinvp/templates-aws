using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Rds.Instance("example", new Aws.Rds.InstanceArgs
        {
            AllocatedStorage = 50,
            MaxAllocatedStorage = 100,
        });
    }

}

