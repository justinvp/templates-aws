import pulumi
import pulumi_aws as aws

example = aws.route53.HealthCheck("example",
    failure_threshold="5",
    fqdn="example.com",
    port=80,
    request_interval="30",
    resource_path="/",
    tags={
        "Name": "tf-test-health-check",
    },
    type="HTTP")

