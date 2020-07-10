import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.efs.AccessPoint("test", {
    fileSystemId: aws_efs_file_system_foo.id,
});

