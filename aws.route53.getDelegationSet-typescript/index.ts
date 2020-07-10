import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dset = pulumi.output(aws.route53.getDelegationSet({
    id: "MQWGHCBFAKEID",
}, { async: true }));

