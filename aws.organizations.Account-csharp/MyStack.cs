using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var account = new Aws.Organizations.Account("account", new Aws.Organizations.AccountArgs
        {
            Email = "john@doe.org",
        });
    }

}

