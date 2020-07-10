import pulumi
import pulumi_aws as aws

front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
front_end_target_group = aws.lb.TargetGroup("frontEndTargetGroup")
front_end_listener = aws.lb.Listener("frontEndListener",
    certificate_arn="arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
    default_actions=[{
        "target_group_arn": front_end_target_group.arn,
        "type": "forward",
    }],
    load_balancer_arn=front_end_load_balancer.arn,
    port="443",
    protocol="HTTPS",
    ssl_policy="ELBSecurityPolicy-2016-08")

