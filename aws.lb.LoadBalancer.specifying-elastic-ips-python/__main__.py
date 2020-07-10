import pulumi
import pulumi_aws as aws

example = aws.lb.LoadBalancer("example",
    load_balancer_type="network",
    subnet_mappings=[
        {
            "allocation_id": aws_eip["example1"]["id"],
            "subnet_id": aws_subnet["example1"]["id"],
        },
        {
            "allocation_id": aws_eip["example2"]["id"],
            "subnet_id": aws_subnet["example2"]["id"],
        },
    ])

