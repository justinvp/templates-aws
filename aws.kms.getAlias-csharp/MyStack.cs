using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var s3 = Output.Create(Aws.Kms.GetAlias.InvokeAsync(new Aws.Kms.GetAliasArgs
        {
            Name = "alias/aws/s3",
        }));
    }

}

