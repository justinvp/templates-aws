import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byteSet = new aws.wafregional.ByteMatchSet("byte_set", {
    byteMatchTuples: [{
        fieldToMatch: {
            data: "referer",
            type: "HEADER",
        },
        positionalConstraint: "CONTAINS",
        targetString: "badrefer1",
        textTransformation: "NONE",
    }],
});

