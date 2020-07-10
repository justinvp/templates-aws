import pulumi
import pulumi_aws as aws

example_load_balancer = aws.lb.LoadBalancer("exampleLoadBalancer",
    internal=True,
    load_balancer_type="network",
    subnet_mappings=[{
        "subnet_id": "12345",
    }])
example_vpc_link = aws.apigateway.VpcLink("exampleVpcLink",
    description="example description",
    target_arn=example_load_balancer.arn)

