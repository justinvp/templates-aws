import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3 = pulumi.output(aws.kms.getAlias({
    name: "alias/aws/s3",
}, { async: true }));

