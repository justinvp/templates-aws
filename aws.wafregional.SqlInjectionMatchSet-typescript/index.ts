import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sqlInjectionMatchSet = new aws.wafregional.SqlInjectionMatchSet("sql_injection_match_set", {
    sqlInjectionMatchTuples: [{
        fieldToMatch: {
            type: "QUERY_STRING",
        },
        textTransformation: "URL_DECODE",
    }],
});

