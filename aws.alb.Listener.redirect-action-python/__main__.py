import pulumi
import pulumi_aws as aws

front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
front_end_listener = aws.lb.Listener("frontEndListener",
    default_actions=[{
        "redirect": {
            "port": "443",
            "protocol": "HTTPS",
            "status_code": "HTTP_301",
        },
        "type": "redirect",
    }],
    load_balancer_arn=front_end_load_balancer.arn,
    port="80",
    protocol="HTTP")

