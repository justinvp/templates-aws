using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var token = Output.Create(Aws.Ecr.GetAuthorizationToken.InvokeAsync());
    }

}

