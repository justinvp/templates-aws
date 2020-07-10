import pulumi
import pulumi_aws as aws

example = aws.datasync.EfsLocation("example",
    ec2_config={
        "securityGroupArns": [aws_security_group["example"]["arn"]],
        "subnetArn": aws_subnet["example"]["arn"],
    },
    efs_file_system_arn=aws_efs_mount_target["example"]["file_system_arn"])

