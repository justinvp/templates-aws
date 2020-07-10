import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fs = new aws.efs.FileSystem("fs", {});
const policy = new aws.efs.FileSystemPolicy("policy", {
    fileSystemId: fs.id,
    policy: pulumi.interpolate`{
    "Version": "2012-10-17",
    "Id": "ExamplePolicy01",
    "Statement": [
        {
            "Sid": "ExampleSatement01",
            "Effect": "Allow",
            "Principal": {
                "AWS": "*"
            },
            "Resource": "${aws_efs_file_system_test.arn}",
            "Action": [
                "elasticfilesystem:ClientMount",
                "elasticfilesystem:ClientWrite"
            ],
            "Condition": {
                "Bool": {
                    "aws:SecureTransport": "true"
                }
            }
        }
    ]
}
`,
});

