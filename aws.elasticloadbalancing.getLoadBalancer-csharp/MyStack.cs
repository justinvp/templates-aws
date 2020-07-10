using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var lbName = config.Get("lbName") ?? "";
        var test = Output.Create(Aws.Elb.GetLoadBalancer.InvokeAsync(new Aws.Elb.GetLoadBalancerArgs
        {
            Name = lbName,
        }));
    }

}

