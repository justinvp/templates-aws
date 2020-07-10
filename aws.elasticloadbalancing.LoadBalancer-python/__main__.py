import pulumi
import pulumi_aws as aws

# Create a new load balancer
bar = aws.elb.LoadBalancer("bar",
    access_logs={
        "bucket": "foo",
        "bucket_prefix": "bar",
        "interval": 60,
    },
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    connection_draining=True,
    connection_draining_timeout=400,
    cross_zone_load_balancing=True,
    health_check={
        "healthyThreshold": 2,
        "interval": 30,
        "target": "HTTP:8000/",
        "timeout": 3,
        "unhealthyThreshold": 2,
    },
    idle_timeout=400,
    instances=[aws_instance["foo"]["id"]],
    listeners=[
        {
            "instance_port": 8000,
            "instanceProtocol": "http",
            "lb_port": 80,
            "lbProtocol": "http",
        },
        {
            "instance_port": 8000,
            "instanceProtocol": "http",
            "lb_port": 443,
            "lbProtocol": "https",
            "sslCertificateId": "arn:aws:iam::123456789012:server-certificate/certName",
        },
    ],
    tags={
        "Name": "foobar-elb",
    })

