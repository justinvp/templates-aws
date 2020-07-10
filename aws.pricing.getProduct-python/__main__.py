import pulumi
import pulumi_aws as aws

example = aws.pricing.get_product(filters=[
        {
            "field": "instanceType",
            "value": "c5.xlarge",
        },
        {
            "field": "operatingSystem",
            "value": "Linux",
        },
        {
            "field": "location",
            "value": "US East (N. Virginia)",
        },
        {
            "field": "preInstalledSw",
            "value": "NA",
        },
        {
            "field": "licenseModel",
            "value": "No License required",
        },
        {
            "field": "tenancy",
            "value": "Shared",
        },
        {
            "field": "capacitystatus",
            "value": "Used",
        },
    ],
    service_code="AmazonEC2")

