import pulumi
import pulumi_aws as aws

network = aws.cloudformation.get_stack(name="my-network-stack")
web = aws.ec2.Instance("web",
    ami="ami-abb07bcb",
    instance_type="t1.micro",
    subnet_id=network.outputs["SubnetId"],
    tags={
        "Name": "HelloWorld",
    })

