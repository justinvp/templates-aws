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
        var pool = new Aws.Cognito.UserPool("pool", new Aws.Cognito.UserPoolArgs
        {
        });
        var client = new Aws.Cognito.UserPoolClient("client", new Aws.Cognito.UserPoolClientArgs
        {
        });
        var domain = new Aws.Cognito.UserPoolDomain("domain", new Aws.Cognito.UserPoolDomainArgs
        {
        });
        var frontEndListener = new Aws.LB.Listener("frontEndListener", new Aws.LB.ListenerArgs
        {
            DefaultActions = 
            {
                new Aws.LB.Inputs.ListenerDefaultActionArgs
                {
                    AuthenticateCognito = new Aws.LB.Inputs.ListenerDefaultActionAuthenticateCognitoArgs
                    {
                        UserPoolArn = pool.Arn,
                        UserPoolClientId = client.Id,
                        UserPoolDomain = domain.Domain,
                    },
                    Type = "authenticate-cognito",
                },
                new Aws.LB.Inputs.ListenerDefaultActionArgs
                {
                    TargetGroupArn = frontEndTargetGroup.Arn,
                    Type = "forward",
                },
            },
            LoadBalancerArn = frontEndLoadBalancer.Arn,
            Port = 80,
            Protocol = "HTTP",
        });
    }

}

