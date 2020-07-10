using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ses.EmailIdentity("example", new Aws.Ses.EmailIdentityArgs
        {
            Email = "email@example.com",
        });
    }

}

