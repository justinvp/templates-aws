import pulumi
import pulumi_aws as aws

ebs_volume = aws.ebs.get_volume(filters=[
        {
            "name": "volume-type",
            "values": ["gp2"],
        },
        {
            "name": "tag:Name",
            "values": ["Example"],
        },
    ],
    most_recent=True)

