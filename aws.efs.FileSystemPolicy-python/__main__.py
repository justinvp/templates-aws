import pulumi
import pulumi_aws as aws

fs = aws.efs.FileSystem("fs")
policy = aws.efs.FileSystemPolicy("policy",
    file_system_id=fs.id,
    policy=f"""{{
    "Version": "2012-10-17",
    "Id": "ExamplePolicy01",
    "Statement": [
        {{
            "Sid": "ExampleSatement01",
            "Effect": "Allow",
            "Principal": {{
                "AWS": "*"
            }},
            "Resource": "{aws_efs_file_system["test"]["arn"]}",
            "Action": [
                "elasticfilesystem:ClientMount",
                "elasticfilesystem:ClientWrite"
            ],
            "Condition": {{
                "Bool": {{
                    "aws:SecureTransport": "true"
                }}
            }}
        }}
    ]
}}

""")

