import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafv2.RegexPatternSet("example", {
    description: "Example regex pattern set",
    regularExpressions: [
        {
            regexString: "one",
        },
        {
            regexString: "two",
        },
    ],
    scope: "REGIONAL",
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
});

