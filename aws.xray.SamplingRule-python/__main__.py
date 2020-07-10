import pulumi
import pulumi_aws as aws

example = aws.xray.SamplingRule("example",
    attributes={
        "Hello": "Tris",
    },
    fixed_rate=0.05,
    host="*",
    http_method="*",
    priority=10000,
    reservoir_size=1,
    resource_arn="*",
    rule_name="example",
    service_name="*",
    service_type="*",
    url_path="*",
    version=1)

