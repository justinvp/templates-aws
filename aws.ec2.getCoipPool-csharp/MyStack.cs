using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var coipPoolId = config.RequireObject<dynamic>("coipPoolId");
        var selected = Output.Create(Aws.Ec2.GetCoipPool.InvokeAsync(new Aws.Ec2.GetCoipPoolArgs
        {
            Id = coipPoolId,
        }));
    }

}

