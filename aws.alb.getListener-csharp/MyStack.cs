using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var listenerArn = config.RequireObject<dynamic>("listenerArn");
        var listener = Output.Create(Aws.LB.GetListener.InvokeAsync(new Aws.LB.GetListenerArgs
        {
            Arn = listenerArn,
        }));
        var selected = Output.Create(Aws.LB.GetLoadBalancer.InvokeAsync(new Aws.LB.GetLoadBalancerArgs
        {
            Name = "default-public",
        }));
        var selected443 = selected.Apply(selected => Output.Create(Aws.LB.GetListener.InvokeAsync(new Aws.LB.GetListenerArgs
        {
            LoadBalancerArn = selected.Arn,
            Port = 443,
        })));
    }

}

