using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var production = Output.Create(Aws.Lambda.GetAlias.InvokeAsync(new Aws.Lambda.GetAliasArgs
        {
            FunctionName = "my-lambda-func",
            Name = "production",
        }));
    }

}

