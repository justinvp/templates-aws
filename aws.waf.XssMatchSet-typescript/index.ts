import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const xssMatchSet = new aws.waf.XssMatchSet("xss_match_set", {
    xssMatchTuples: [
        {
            fieldToMatch: {
                type: "URI",
            },
            textTransformation: "NONE",
        },
        {
            fieldToMatch: {
                type: "QUERY_STRING",
            },
            textTransformation: "NONE",
        },
    ],
});

