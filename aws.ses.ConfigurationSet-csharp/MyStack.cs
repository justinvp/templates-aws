using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.Ses.ConfigurationSet("test", new Aws.Ses.ConfigurationSetArgs
        {
        });
    }

}

