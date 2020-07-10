import pulumi
import pulumi_aws as aws

example = aws.get_ami(executable_users=["self"],
    filters=[
        {
            "name": "name",
            "values": ["myami-*"],
        },
        {
            "name": "root-device-type",
            "values": ["ebs"],
        },
        {
            "name": "virtualization-type",
            "values": ["hvm"],
        },
    ],
    most_recent=True,
    name_regex="^myami-\\d{3}",
    owners=["self"])

