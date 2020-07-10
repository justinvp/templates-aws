import pulumi
import pulumi_aws as aws

front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
front_end_target_group = aws.lb.TargetGroup("frontEndTargetGroup")
pool = aws.cognito.UserPool("pool")
client = aws.cognito.UserPoolClient("client")
domain = aws.cognito.UserPoolDomain("domain")
front_end_listener = aws.lb.Listener("frontEndListener",
    default_actions=[
        {
            "authenticateCognito": {
                "userPoolArn": pool.arn,
                "userPoolClientId": client.id,
                "userPoolDomain": domain.domain,
            },
            "type": "authenticate-cognito",
        },
        {
            "target_group_arn": front_end_target_group.arn,
            "type": "forward",
        },
    ],
    load_balancer_arn=front_end_load_balancer.arn,
    port="80",
    protocol="HTTP")

