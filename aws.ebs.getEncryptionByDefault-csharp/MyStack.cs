using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.Ebs.GetEncryptionByDefault.InvokeAsync());
    }

}

