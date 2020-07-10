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
            DefaultActions = 
            {
                new Aws.LB.Inputs.ListenerDefaultActionArgs
                {
                    AuthenticateOidc = new Aws.LB.Inputs.ListenerDefaultActionAuthenticateOidcArgs
                    {
                        AuthorizationEndpoint = "https://example.com/authorization_endpoint",
                        ClientId = "client_id",
                        ClientSecret = "client_secret",
                        Issuer = "https://example.com",
                        TokenEndpoint = "https://example.com/token_endpoint",
                        UserInfoEndpoint = "https://example.com/user_info_endpoint",
                    },
                    Type = "authenticate-oidc",
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

