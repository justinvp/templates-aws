import pulumi
import pulumi_aws as aws

web = aws.ec2.Instance("web",
    ami="ami-21f78e11",
    availability_zone="us-west-2a",
    instance_type="t1.micro",
    tags={
        "Name": "HelloWorld",
    })
example = aws.ec2.Eip("example", vpc=True)
eip_assoc = aws.ec2.EipAssociation("eipAssoc",
    allocation_id=example.id,
    instance_id=web.id)

