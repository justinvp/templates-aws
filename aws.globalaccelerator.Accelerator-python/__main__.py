import pulumi
import pulumi_aws as aws

example = aws.globalaccelerator.Accelerator("example",
    attributes={
        "flowLogsEnabled": True,
        "flowLogsS3Bucket": "example-bucket",
        "flowLogsS3Prefix": "flow-logs/",
    },
    enabled=True,
    ip_address_type="IPV4")

