import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.fsx.WindowsFileSystem("example", {
    activeDirectoryId: aws_directory_service_directory_example.id,
    kmsKeyId: aws_kms_key_example.arn,
    storageCapacity: 300,
    subnetIds: aws_subnet_example.id,
    throughputCapacity: 1024,
});

