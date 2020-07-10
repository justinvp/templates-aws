import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc",
    cidr_block="10.0.0.0/16",
    enable_dns_hostnames=True,
    enable_dns_support=True)
example_private_dns_namespace = aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace",
    description="example",
    vpc=example_vpc.id)
example_service = aws.servicediscovery.Service("exampleService",
    dns_config={
        "dnsRecords": [{
            "ttl": 10,
            "type": "A",
        }],
        "namespace_id": example_private_dns_namespace.id,
        "routingPolicy": "MULTIVALUE",
    },
    health_check_custom_config={
        "failure_threshold": 1,
    })

