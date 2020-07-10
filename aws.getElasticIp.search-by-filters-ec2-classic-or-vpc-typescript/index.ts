import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byFilter = pulumi.output(aws.getElasticIp({
    filters: [{
        name: "tag:Name",
        values: ["exampleNameTagValue"],
    }],
}, { async: true }));

