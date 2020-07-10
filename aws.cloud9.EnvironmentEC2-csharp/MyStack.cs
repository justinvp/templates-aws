using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Cloud9.EnvironmentEC2("example", new Aws.Cloud9.EnvironmentEC2Args
        {
            InstanceType = "t2.micro",
        });
    }

}

