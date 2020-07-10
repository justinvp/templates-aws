import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRule = new aws.waf.Rule("example", {
    metricName: "example",
});
const exampleRuleGroup = new aws.waf.RuleGroup("example", {
    activatedRules: [{
        action: {
            type: "COUNT",
        },
        priority: 50,
        ruleId: exampleRule.id,
    }],
    metricName: "example",
});

