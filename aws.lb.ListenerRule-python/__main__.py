import pulumi
import pulumi_aws as aws

front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
front_end_listener = aws.lb.Listener("frontEndListener")
static = aws.lb.ListenerRule("static",
    actions=[{
        "target_group_arn": aws_lb_target_group["static"]["arn"],
        "type": "forward",
    }],
    conditions=[
        {
            "pathPattern": {
                "values": ["/static/*"],
            },
        },
        {
            "hostHeader": {
                "values": ["example.com"],
            },
        },
    ],
    listener_arn=front_end_listener.arn,
    priority=100)
host_based_routing = aws.lb.ListenerRule("hostBasedRouting",
    actions=[{
        "forward": {
            "stickiness": {
                "duration": 600,
                "enabled": True,
            },
            "targetGroup": [
                {
                    "arn": aws_lb_target_group["main"]["arn"],
                    "weight": 80,
                },
                {
                    "arn": aws_lb_target_group["canary"]["arn"],
                    "weight": 20,
                },
            ],
        },
        "type": "forward",
    }],
    conditions=[{
        "hostHeader": {
            "values": ["my-service.*.mycompany.io"],
        },
    }],
    listener_arn=front_end_listener.arn,
    priority=99)
host_based_weighted_routing = aws.lb.ListenerRule("hostBasedWeightedRouting",
    actions=[{
        "target_group_arn": aws_lb_target_group["static"]["arn"],
        "type": "forward",
    }],
    conditions=[{
        "hostHeader": {
            "values": ["my-service.*.mydomain.io"],
        },
    }],
    listener_arn=front_end_listener.arn,
    priority=99)
redirect_http_to_https = aws.lb.ListenerRule("redirectHttpToHttps",
    actions=[{
        "redirect": {
            "port": "443",
            "protocol": "HTTPS",
            "status_code": "HTTP_301",
        },
        "type": "redirect",
    }],
    conditions=[{
        "httpHeader": {
            "httpHeaderName": "X-Forwarded-For",
            "values": ["192.168.1.*"],
        },
    }],
    listener_arn=front_end_listener.arn)
health_check = aws.lb.ListenerRule("healthCheck",
    actions=[{
        "fixedResponse": {
            "content_type": "text/plain",
            "messageBody": "HEALTHY",
            "status_code": "200",
        },
        "type": "fixed-response",
    }],
    conditions=[{
        "queryString": [
            {
                "key": "health",
                "value": "check",
            },
            {
                "value": "bar",
            },
        ],
    }],
    listener_arn=front_end_listener.arn)
pool = aws.cognito.UserPool("pool")
client = aws.cognito.UserPoolClient("client")
domain = aws.cognito.UserPoolDomain("domain")
admin = aws.lb.ListenerRule("admin",
    actions=[
        {
            "authenticateOidc": {
                "authorizationEndpoint": "https://example.com/authorization_endpoint",
                "client_id": "client_id",
                "client_secret": "client_secret",
                "issuer": "https://example.com",
                "tokenEndpoint": "https://example.com/token_endpoint",
                "userInfoEndpoint": "https://example.com/user_info_endpoint",
            },
            "type": "authenticate-oidc",
        },
        {
            "target_group_arn": aws_lb_target_group["static"]["arn"],
            "type": "forward",
        },
    ],
    listener_arn=front_end_listener.arn)

