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
                    InstancePort = 25,
                    InstanceProtocol = "tcp",
                    LbPort = 25,
                    LbProtocol = "tcp",
                },
                new Aws.Elb.Inputs.LoadBalancerListenerArgs
                {
                    InstancePort = 587,
                    InstanceProtocol = "tcp",
                    LbPort = 587,
                    LbProtocol = "tcp",
                },
            },
        });
        var smtp = new Aws.Ec2.ProxyProtocolPolicy("smtp", new Aws.Ec2.ProxyProtocolPolicyArgs
        {
            InstancePorts = 
            {
                "25",
                "587",
            },
            LoadBalancer = lb.Name,
        });
    }

}

