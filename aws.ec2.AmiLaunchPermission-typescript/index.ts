import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.AmiLaunchPermission("example", {
    accountId: "123456789012",
    imageId: "ami-12345678",
});

