import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const stream = pulumi.output(aws.kinesis.getStream({
    name: "stream-name",
}, { async: true }));

