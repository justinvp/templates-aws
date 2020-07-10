import pulumi
import pulumi_aws as aws

example = aws.route53.HealthCheck("example",
    failure_threshold="5",
    fqdn="example.com",
    port=443,
    request_interval="30",
    resource_path="/",
    search_string="example",
    type="HTTPS_STR_MATCH")

