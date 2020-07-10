using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ssm.Parameter("foo", new Aws.Ssm.ParameterArgs
        {
            Type = "String",
            Value = "bar",
        });
    }

}

