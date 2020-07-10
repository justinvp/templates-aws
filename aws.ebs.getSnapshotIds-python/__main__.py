import pulumi
import pulumi_aws as aws

ebs_volumes = aws.ebs.get_snapshot_ids(filters=[
        {
            "name": "volume-size",
            "values": ["40"],
        },
        {
            "name": "tag:Name",
            "values": ["Example"],
        },
    ],
    owners=["self"])

