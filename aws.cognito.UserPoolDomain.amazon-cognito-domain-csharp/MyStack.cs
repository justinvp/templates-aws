using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Cognito.UserPool("example", new Aws.Cognito.UserPoolArgs
        {
        });
        var main = new Aws.Cognito.UserPoolDomain("main", new Aws.Cognito.UserPoolDomainArgs
        {
            Domain = "example-domain",
            UserPoolId = example.Id,
        });
    }

}

