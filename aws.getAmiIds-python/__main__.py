import pulumi
import pulumi_aws as aws

ubuntu = aws.get_ami_ids(filters=[{
        "name": "name",
        "values": ["ubuntu/images/ubuntu-*-*-amd64-server-*"],
    }],
    owners=["099720109477"])

