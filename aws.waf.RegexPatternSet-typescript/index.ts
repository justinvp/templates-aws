import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.waf.RegexPatternSet("example", {
    regexPatternStrings: [
        "one",
        "two",
    ],
});

