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
                    FixedResponse = new Aws.LB.Inputs.ListenerDefaultActionFixedResponseArgs
                    {
                        ContentType = "text/plain",
                        MessageBody = "Fixed response content",
                        StatusCode = "200",
                    },
                    Type = "fixed-response",
                },
            },
            LoadBalancerArn = frontEndLoadBalancer.Arn,
            Port = 80,
            Protocol = "HTTP",
        });
    }

}

