import pulumi
import pulumi_aws as aws

example = aws.licensemanager.LicenseConfiguration("example",
    description="Example",
    license_count=10,
    license_count_hard_limit=True,
    license_counting_type="Socket",
    license_rules=["#minimumSockets=2"],
    tags={
        "foo": "barr",
    })

