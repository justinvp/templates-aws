import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafregional.RegexPatternSet("example", {
    regexPatternStrings: [
        "one",
        "two",
    ],
});

