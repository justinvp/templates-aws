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
        });
        var @static = new Aws.LB.ListenerRule("static", new Aws.LB.ListenerRuleArgs
        {
            Actions = 
            {
                new Aws.LB.Inputs.ListenerRuleActionArgs
                {
                    TargetGroupArn = aws_lb_target_group.Static.Arn,
                    Type = "forward",
                },
            },
            Conditions = 
            {
                new Aws.LB.Inputs.ListenerRuleConditionArgs
                {
                    PathPattern = new Aws.LB.Inputs.ListenerRuleConditionPathPatternArgs
                    {
                        Values = 
                        {
                            "/static/*",
                        },
                    },
                },
                new Aws.LB.Inputs.ListenerRuleConditionArgs
                {
                    HostHeader = new Aws.LB.Inputs.ListenerRuleConditionHostHeaderArgs
                    {
                        Values = 
                        {
                            "example.com",
                        },
                    },
                },
            },
            ListenerArn = frontEndListener.Arn,
            Priority = 100,
        });
        var hostBasedRouting = new Aws.LB.ListenerRule("hostBasedRouting", new Aws.LB.ListenerRuleArgs
        {
            Actions = 
            {
                new Aws.LB.Inputs.ListenerRuleActionArgs
                {
                    Forward = new Aws.LB.Inputs.ListenerRuleActionForwardArgs
                    {
                        Stickiness = new Aws.LB.Inputs.ListenerRuleActionForwardStickinessArgs
                        {
                            Duration = 600,
                            Enabled = true,
                        },
                        TargetGroup = 
                        {
                            
                            {
                                { "arn", aws_lb_target_group.Main.Arn },
                                { "weight", 80 },
                            },
                            
                            {
                                { "arn", aws_lb_target_group.Canary.Arn },
                                { "weight", 20 },
                            },
                        },
                    },
                    Type = "forward",
                },
            },
            Conditions = 
            {
                new Aws.LB.Inputs.ListenerRuleConditionArgs
                {
                    HostHeader = new Aws.LB.Inputs.ListenerRuleConditionHostHeaderArgs
                    {
                        Values = 
                        {
                            "my-service.*.mycompany.io",
                        },
                    },
                },
            },
            ListenerArn = frontEndListener.Arn,
            Priority = 99,
        });
        var hostBasedWeightedRouting = new Aws.LB.ListenerRule("hostBasedWeightedRouting", new Aws.LB.ListenerRuleArgs
        {
            Actions = 
            {
                new Aws.LB.Inputs.ListenerRuleActionArgs
                {
                    TargetGroupArn = aws_lb_target_group.Static.Arn,
                    Type = "forward",
                },
            },
            Conditions = 
            {
                new Aws.LB.Inputs.ListenerRuleConditionArgs
                {
                    HostHeader = new Aws.LB.Inputs.ListenerRuleConditionHostHeaderArgs
                    {
                        Values = 
                        {
                            "my-service.*.mydomain.io",
                        },
                    },
                },
            },
            ListenerArn = frontEndListener.Arn,
            Priority = 99,
        });
        var redirectHttpToHttps = new Aws.LB.ListenerRule("redirectHttpToHttps", new Aws.LB.ListenerRuleArgs
        {
            Actions = 
            {
                new Aws.LB.Inputs.ListenerRuleActionArgs
                {
                    Redirect = new Aws.LB.Inputs.ListenerRuleActionRedirectArgs
                    {
                        Port = "443",
                        Protocol = "HTTPS",
                        StatusCode = "HTTP_301",
                    },
                    Type = "redirect",
                },
            },
            Conditions = 
            {
                new Aws.LB.Inputs.ListenerRuleConditionArgs
                {
                    HttpHeader = new Aws.LB.Inputs.ListenerRuleConditionHttpHeaderArgs
                    {
                        HttpHeaderName = "X-Forwarded-For",
                        Values = 
                        {
                            "192.168.1.*",
                        },
                    },
                },
            },
            ListenerArn = frontEndListener.Arn,
        });
        var healthCheck = new Aws.LB.ListenerRule("healthCheck", new Aws.LB.ListenerRuleArgs
        {
            Actions = 
            {
                new Aws.LB.Inputs.ListenerRuleActionArgs
                {
                    FixedResponse = new Aws.LB.Inputs.ListenerRuleActionFixedResponseArgs
                    {
                        ContentType = "text/plain",
                        MessageBody = "HEALTHY",
                        StatusCode = "200",
                    },
                    Type = "fixed-response",
                },
            },
            Conditions = 
            {
                new Aws.LB.Inputs.ListenerRuleConditionArgs
                {
                    QueryString = 
                    {
                        
                        {
                            { "key", "health" },
                            { "value", "check" },
                        },
                        
                        {
                            { "value", "bar" },
                        },
                    },
                },
            },
            ListenerArn = frontEndListener.Arn,
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
        var admin = new Aws.LB.ListenerRule("admin", new Aws.LB.ListenerRuleArgs
        {
            Actions = 
            {
                new Aws.LB.Inputs.ListenerRuleActionArgs
                {
                    AuthenticateOidc = new Aws.LB.Inputs.ListenerRuleActionAuthenticateOidcArgs
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
                new Aws.LB.Inputs.ListenerRuleActionArgs
                {
                    TargetGroupArn = aws_lb_target_group.Static.Arn,
                    Type = "forward",
                },
            },
            ListenerArn = frontEndListener.Arn,
        });
    }

}

