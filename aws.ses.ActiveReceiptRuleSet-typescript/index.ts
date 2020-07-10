import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ses.ActiveReceiptRuleSet("main", {
    ruleSetName: "primary-rules",
});

