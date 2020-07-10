import pulumi
import pulumi_aws as aws

example = aws.get_availability_zones(all_availability_zones=True,
    filters=[{
        "name": "opt-in-status",
        "values": [
            "not-opted-in",
            "opted-in",
        ],
    }])

