import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byTags = pulumi.output(aws.getElasticIp({
    tags: {
        Name: "exampleNameTagValue",
    },
}, { async: true }));

