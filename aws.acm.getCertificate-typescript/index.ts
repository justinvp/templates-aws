import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Find a RSA 4096 bit certificate
const example = pulumi.output(aws.acm.getCertificate({
    domain: "tf.example.com",
    keyTypes: ["RSA_4096"],
}, { async: true }));

