using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var pool = new Aws.Cognito.UserPool("pool", new Aws.Cognito.UserPoolArgs
        {
        });
    }

}

