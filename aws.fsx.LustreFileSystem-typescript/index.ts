import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.fsx.LustreFileSystem("example", {
    importPath: pulumi.interpolate`s3://${aws_s3_bucket_example.bucket}`,
    storageCapacity: 1200,
    subnetIds: aws_subnet_example.id,
});

