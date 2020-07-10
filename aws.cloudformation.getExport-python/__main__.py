import pulumi
import pulumi_aws as aws

subnet_id = aws.cloudformation.get_export(name="mySubnetIdExportName")
web = aws.ec2.Instance("web",
    ami="ami-abb07bcb",
    instance_type="t1.micro",
    subnet_id=subnet_id.value)

