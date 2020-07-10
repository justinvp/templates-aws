import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byPublicIp = pulumi.output(aws.getElasticIp({
    publicIp: "1.2.3.4",
}, { async: true }));

