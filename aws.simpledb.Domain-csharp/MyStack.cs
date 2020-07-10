using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var users = new Aws.SimpleDB.Domain("users", new Aws.SimpleDB.DomainArgs
        {
        });
    }

}

