import pulumi
import pulumi_aws as aws

front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
front_end_target_group = aws.lb.TargetGroup("frontEndTargetGroup")
front_end_listener = aws.lb.Listener("frontEndListener",
    default_actions=[
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
            "target_group_arn": front_end_target_group.arn,
            "type": "forward",
        },
    ],
    load_balancer_arn=front_end_load_balancer.arn,
    port="80",
    protocol="HTTP")

