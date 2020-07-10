import pulumi
import pulumi_aws as aws

test = aws.lb.LoadBalancer("test",
    access_logs={
        "bucket": aws_s3_bucket["lb_logs"]["bucket"],
        "enabled": True,
        "prefix": "test-lb",
    },
    enable_deletion_protection=True,
    internal=False,
    load_balancer_type="application",
    security_groups=[aws_security_group["lb_sg"]["id"]],
    subnets=[[__item["id"] for __item in aws_subnet["public"]]],
    tags={
        "Environment": "production",
    })

