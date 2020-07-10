import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.getAvailabilityZones({
    allAvailabilityZones: true,
    filters: [{
        name: "opt-in-status",
        values: [
            "not-opted-in",
            "opted-in",
        ],
    }],
}, { async: true }));

