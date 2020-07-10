import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selected = pulumi.output(aws.s3.getBucket({
    bucket: "a-test-bucket",
}, { async: true }));
const test = new aws.cloudfront.Distribution("test", {
    origins: [{
        domainName: selected.bucketDomainName,
        originId: "s3-selected-bucket",
    }],
});

