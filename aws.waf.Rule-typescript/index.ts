import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ipset = new aws.waf.IpSet("ipset", {
    ipSetDescriptors: [{
        type: "IPV4",
        value: "192.0.7.0/24",
    }],
});
const wafrule = new aws.waf.Rule("wafrule", {
    metricName: "tfWAFRule",
    predicates: [{
        dataId: ipset.id,
        negated: false,
        type: "IPMatch",
    }],
}, { dependsOn: [ipset] });

