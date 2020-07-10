using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.SecurityHub.Account("example", new Aws.SecurityHub.AccountArgs
        {
        });
    }

}

