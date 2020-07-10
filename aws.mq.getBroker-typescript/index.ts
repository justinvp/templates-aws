import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const brokerId = config.get("brokerId") || "";
const brokerName = config.get("brokerName") || "";

const byId = pulumi.output(aws.mq.getBroker({
    brokerId: brokerId,
}, { async: true }));
const byName = pulumi.output(aws.mq.getBroker({
    brokerName: brokerName,
}, { async: true }));

