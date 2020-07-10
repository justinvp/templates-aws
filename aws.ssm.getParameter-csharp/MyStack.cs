using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = Output.Create(Aws.Ssm.GetParameter.InvokeAsync(new Aws.Ssm.GetParameterArgs
        {
            Name = "foo",
        }));
    }

}

