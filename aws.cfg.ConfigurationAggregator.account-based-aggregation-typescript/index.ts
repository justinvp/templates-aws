import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const account = new aws.cfg.ConfigurationAggregator("account", {
    accountAggregationSource: {
        accountIds: ["123456789012"],
        regions: ["us-west-2"],
    },
});

