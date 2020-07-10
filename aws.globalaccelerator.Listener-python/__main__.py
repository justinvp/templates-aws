import pulumi
import pulumi_aws as aws

example_accelerator = aws.globalaccelerator.Accelerator("exampleAccelerator",
    attributes={
        "flowLogsEnabled": True,
        "flowLogsS3Bucket": "example-bucket",
        "flowLogsS3Prefix": "flow-logs/",
    },
    enabled=True,
    ip_address_type="IPV4")
example_listener = aws.globalaccelerator.Listener("exampleListener",
    accelerator_arn=example_accelerator.id,
    client_affinity="SOURCE_IP",
    port_ranges=[{
        "from_port": 80,
        "to_port": 80,
    }],
    protocol="TCP")

