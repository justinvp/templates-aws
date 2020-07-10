import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byAllocationId = pulumi.output(aws.getElasticIp({
    id: "eipalloc-12345678",
}, { async: true }));

