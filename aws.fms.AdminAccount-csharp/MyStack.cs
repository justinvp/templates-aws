using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Fms.AdminAccount("example", new Aws.Fms.AdminAccountArgs
        {
        });
    }

}

