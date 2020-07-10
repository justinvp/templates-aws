import pulumi
import pulumi_aws as aws

example = aws.fsx.WindowsFileSystem("example",
    kms_key_id=aws_kms_key["example"]["arn"],
    self_managed_active_directory={
        "dns_ips": [
            "10.0.0.111",
            "10.0.0.222",
        ],
        "domain_name": "corp.example.com",
        "password": "avoid-plaintext-passwords",
        "username": "Admin",
    },
    storage_capacity=300,
    subnet_ids=aws_subnet["example"]["id"],
    throughput_capacity=1024)

