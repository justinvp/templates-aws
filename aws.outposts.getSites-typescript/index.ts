import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const all = pulumi.output(aws.outposts.getSites({ async: true }));

