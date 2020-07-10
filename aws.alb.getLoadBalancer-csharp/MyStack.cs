using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var lbArn = config.Get("lbArn") ?? "";
        var lbName = config.Get("lbName") ?? "";
        var test = Output.Create(Aws.LB.GetLoadBalancer.InvokeAsync(new Aws.LB.GetLoadBalancerArgs
        {
            Arn = lbArn,
            Name = lbName,
        }));
    }

}

