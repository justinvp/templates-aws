import pulumi
import pulumi_aws as aws

example = aws.fsx.WindowsFileSystem("example",
    active_directory_id=aws_directory_service_directory["example"]["id"],
    kms_key_id=aws_kms_key["example"]["arn"],
    storage_capacity=300,
    subnet_ids=aws_subnet["example"]["id"],
    throughput_capacity=1024)

