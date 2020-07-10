using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var frontEndLoadBalancer = new Aws.LB.LoadBalancer("frontEndLoadBalancer", new Aws.LB.LoadBalancerArgs
        {
        });
        var frontEndTargetGroup = new Aws.LB.TargetGroup("frontEndTargetGroup", new Aws.LB.TargetGroupArgs
        {
        });
        var frontEndListener = new Aws.LB.Listener("frontEndListener", new Aws.LB.ListenerArgs
        {
            CertificateArn = "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
            DefaultActions = 
            {
                new Aws.LB.Inputs.ListenerDefaultActionArgs
                {
                    TargetGroupArn = frontEndTargetGroup.Arn,
                    Type = "forward",
                },
            },
            LoadBalancerArn = frontEndLoadBalancer.Arn,
            Port = 443,
            Protocol = "HTTPS",
            SslPolicy = "ELBSecurityPolicy-2016-08",
        });
    }

}

