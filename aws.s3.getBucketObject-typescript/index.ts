import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bootstrapScript = pulumi.output(aws.s3.getBucketObject({
    bucket: "ourcorp-deploy-config",
    key: "ec2-bootstrap-script.sh",
}, { async: true }));
const example = new aws.ec2.Instance("example", {
    ami: "ami-2757f631",
    instanceType: "t2.micro",
    userData: bootstrapScript.body,
});

