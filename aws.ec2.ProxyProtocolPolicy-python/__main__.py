import pulumi
import pulumi_aws as aws

lb = aws.elb.LoadBalancer("lb",
    availability_zones=["us-east-1a"],
    listeners=[
        {
            "instance_port": 25,
            "instanceProtocol": "tcp",
            "lb_port": 25,
            "lbProtocol": "tcp",
        },
        {
            "instance_port": 587,
            "instanceProtocol": "tcp",
            "lb_port": 587,
            "lbProtocol": "tcp",
        },
    ])
smtp = aws.ec2.ProxyProtocolPolicy("smtp",
    instance_ports=[
        "25",
        "587",
    ],
    load_balancer=lb.name)

