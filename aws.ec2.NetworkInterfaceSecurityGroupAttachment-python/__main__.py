import pulumi
import pulumi_aws as aws

ami = aws.get_ami(filters=[{
        "name": "name",
        "values": ["amzn-ami-hvm-*"],
    }],
    most_recent=True,
    owners=["amazon"])
instance = aws.ec2.Instance("instance",
    ami=ami.id,
    instance_type="t2.micro",
    tags={
        "type": "test-instance",
    })
sg = aws.ec2.SecurityGroup("sg", tags={
    "type": "test-security-group",
})
sg_attachment = aws.ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment",
    network_interface_id=instance.primary_network_interface_id,
    security_group_id=sg.id)

