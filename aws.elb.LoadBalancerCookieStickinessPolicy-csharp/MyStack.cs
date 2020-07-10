using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var lb = new Aws.Elb.LoadBalancer("lb", new Aws.Elb.LoadBalancerArgs
        {
            AvailabilityZones = 
            {
                "us-east-1a",
            },
            Listeners = 
            {
                new Aws.Elb.Inputs.LoadBalancerListenerArgs
                {
                    InstancePort = 8000,
                    InstanceProtocol = "http",
                    LbPort = 80,
                    LbProtocol = "http",
                },
            },
        });
        var foo = new Aws.Elb.LoadBalancerCookieStickinessPolicy("foo", new Aws.Elb.LoadBalancerCookieStickinessPolicyArgs
        {
            CookieExpirationPeriod = 600,
            LbPort = 80,
            LoadBalancer = lb.Id,
        });
    }

}

