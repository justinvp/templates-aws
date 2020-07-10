import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRegexPatternSet = new aws.wafregional.RegexPatternSet("example", {
    regexPatternStrings: [
        "one",
        "two",
    ],
});
const exampleRegexMatchSet = new aws.wafregional.RegexMatchSet("example", {
    regexMatchTuples: [{
        fieldToMatch: {
            data: "User-Agent",
            type: "HEADER",
        },
        regexPatternSetId: exampleRegexPatternSet.id,
        textTransformation: "NONE",
    }],
});

