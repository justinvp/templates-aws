using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Cfg.AggregateAuthorization("example", new Aws.Cfg.AggregateAuthorizationArgs
        {
            AccountId = "123456789012",
            Region = "eu-west-2",
        });
    }

}

