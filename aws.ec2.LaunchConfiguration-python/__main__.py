import pulumi
import pulumi_aws as aws

ubuntu = aws.get_ami(filters=[
        {
            "name": "name",
            "values": ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
        },
        {
            "name": "virtualization-type",
            "values": ["hvm"],
        },
    ],
    most_recent=True,
    owners=["099720109477"])
as_conf = aws.ec2.LaunchConfiguration("asConf",
    image_id=ubuntu.id,
    instance_type="t2.micro")

