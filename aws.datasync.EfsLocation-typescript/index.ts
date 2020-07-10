import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.EfsLocation("example", {
    ec2Config: {
        securityGroupArns: [aws_security_group_example.arn],
        subnetArn: aws_subnet_example.arn,
    },
    // The below example uses aws_efs_mount_target as a reference to ensure a mount target already exists when resource creation occurs.
    // You can accomplish the same behavior with depends_on or an aws_efs_mount_target data source reference.
    efsFileSystemArn: aws_efs_mount_target_example.fileSystemArn,
});

