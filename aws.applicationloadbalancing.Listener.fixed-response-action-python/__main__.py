import pulumi
import pulumi_aws as aws

front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
front_end_listener = aws.lb.Listener("frontEndListener",
    default_actions=[{
        "fixedResponse": {
            "content_type": "text/plain",
            "messageBody": "Fixed response content",
            "status_code": "200",
        },
        "type": "fixed-response",
    }],
    load_balancer_arn=front_end_load_balancer.arn,
    port="80",
    protocol="HTTP")

