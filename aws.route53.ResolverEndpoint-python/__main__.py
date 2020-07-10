import pulumi
import pulumi_aws as aws

foo = aws.route53.ResolverEndpoint("foo",
    direction="INBOUND",
    ip_addresses=[
        {
            "subnet_id": aws_subnet["sn1"]["id"],
        },
        {
            "ip": "10.0.64.4",
            "subnet_id": aws_subnet["sn2"]["id"],
        },
    ],
    security_group_ids=[
        aws_security_group["sg1"]["id"],
        aws_security_group["sg2"]["id"],
    ],
    tags={
        "Environment": "Prod",
    })

