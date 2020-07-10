import pulumi
import pulumi_aws as aws

example = aws.directoryservice.ConditionalForwader("example",
    directory_id=aws_directory_service_directory["ad"]["id"],
    dns_ips=[
        "8.8.8.8",
        "8.8.4.4",
    ],
    remote_domain_name="example.com")

