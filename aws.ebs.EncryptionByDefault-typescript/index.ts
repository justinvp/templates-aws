import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.EncryptionByDefault("example", {
    enabled: true,
});

