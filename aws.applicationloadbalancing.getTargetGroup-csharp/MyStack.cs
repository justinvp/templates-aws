using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var lbTgArn = config.Get("lbTgArn") ?? "";
        var lbTgName = config.Get("lbTgName") ?? "";
        var test = Output.Create(Aws.LB.GetTargetGroup.InvokeAsync(new Aws.LB.GetTargetGroupArgs
        {
            Arn = lbTgArn,
            Name = lbTgName,
        }));
    }

}

