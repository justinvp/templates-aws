import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ipset = new aws.wafregional.IpSet("ipset", {
    ipSetDescriptors: [{
        type: "IPV4",
        value: "192.0.7.0/24",
    }],
});
const wafrule = new aws.wafregional.Rule("wafrule", {
    metricName: "tfWAFRule",
    predicates: [{
        dataId: ipset.id,
        negated: false,
        type: "IPMatch",
    }],
});
const wafacl = new aws.wafregional.WebAcl("wafacl", {
    defaultAction: {
        type: "ALLOW",
    },
    metricName: "tfWebACL",
    rules: [{
        action: {
            type: "BLOCK",
        },
        priority: 1,
        ruleId: wafrule.id,
        type: "REGULAR",
    }],
});

