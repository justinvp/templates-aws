import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const layerName = config.require("layerName");

const existing = pulumi.output(aws.lambda.getLayerVersion({
    layerName: layerName,
}, { async: true }));

