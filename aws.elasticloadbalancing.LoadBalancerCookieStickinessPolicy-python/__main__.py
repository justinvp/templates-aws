import pulumi
import pulumi_aws as aws

lb = aws.elb.LoadBalancer("lb",
    availability_zones=["us-east-1a"],
    listeners=[{
        "instance_port": 8000,
        "instanceProtocol": "http",
        "lb_port": 80,
        "lbProtocol": "http",
    }])
foo = aws.elb.LoadBalancerCookieStickinessPolicy("foo",
    cookie_expiration_period=600,
    lb_port=80,
    load_balancer=lb.id)

