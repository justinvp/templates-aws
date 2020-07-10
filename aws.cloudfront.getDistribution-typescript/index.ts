import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.output(aws.cloudfront.getDistribution({
    id: "EDFDVBD632BHDS5",
}, { async: true }));

