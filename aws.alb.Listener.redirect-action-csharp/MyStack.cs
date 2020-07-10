using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var frontEndLoadBalancer = new Aws.LB.LoadBalancer("frontEndLoadBalancer", new Aws.LB.LoadBalancerArgs
        {
        });
        var frontEndListener = new Aws.LB.Listener("frontEndListener", new Aws.LB.ListenerArgs
        {
            DefaultActions = 
            {
                new Aws.LB.Inputs.ListenerDefaultActionArgs
                {
                    Redirect = new Aws.LB.Inputs.ListenerDefaultActionRedirectArgs
                    {
                        Port = "443",
                        Protocol = "HTTPS",
                        StatusCode = "HTTP_301",
                    },
                    Type = "redirect",
                },
            },
            LoadBalancerArn = frontEndLoadBalancer.Arn,
            Port = 80,
            Protocol = "HTTP",
        });
    }

}

