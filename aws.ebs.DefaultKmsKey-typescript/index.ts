import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.DefaultKmsKey("example", {
    keyArn: aws_kms_key_example.arn,
});

