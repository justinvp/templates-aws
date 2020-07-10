import pulumi
import pulumi_aws as aws

ebs_volume = aws.ebs.get_snapshot(filters=[
        {
            "name": "volume-size",
            "values": ["40"],
        },
        {
            "name": "tag:Name",
            "values": ["Example"],
        },
    ],
    most_recent=True,
    owners=["self"])

