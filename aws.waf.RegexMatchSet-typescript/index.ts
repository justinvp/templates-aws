import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRegexPatternSet = new aws.waf.RegexPatternSet("example", {
    regexPatternStrings: [
        "one",
        "two",
    ],
});
const exampleRegexMatchSet = new aws.waf.RegexMatchSet("example", {
    regexMatchTuples: [{
        fieldToMatch: {
            data: "User-Agent",
            type: "HEADER",
        },
        regexPatternSetId: exampleRegexPatternSet.id,
        textTransformation: "NONE",
    }],
});

