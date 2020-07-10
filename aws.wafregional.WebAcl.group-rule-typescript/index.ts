import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafregional.WebAcl("example", {
    defaultAction: {
        type: "ALLOW",
    },
    metricName: "example",
    rules: [{
        overrideAction: {
            type: "NONE",
        },
        priority: 1,
        ruleId: aws_wafregional_rule_group_example.id,
        type: "GROUP",
    }],
});

