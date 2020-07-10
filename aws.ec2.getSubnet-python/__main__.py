import pulumi
import pulumi_aws as aws

config = pulumi.Config()
subnet_id = config.require_object("subnetId")
selected = aws.ec2.get_subnet(id=subnet_id)
subnet = aws.ec2.SecurityGroup("subnet",
    ingress=[{
        "cidr_blocks": [selected.cidr_block],
        "from_port": 80,
        "protocol": "tcp",
        "to_port": 80,
    }],
    vpc_id=selected.vpc_id)

