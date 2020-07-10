import pulumi
import pulumi_aws as aws

test = aws.lb.LoadBalancer("test",
    enable_deletion_protection=True,
    internal=False,
    load_balancer_type="network",
    subnets=[[__item["id"] for __item in aws_subnet["public"]]],
    tags={
        "Environment": "production",
    })

