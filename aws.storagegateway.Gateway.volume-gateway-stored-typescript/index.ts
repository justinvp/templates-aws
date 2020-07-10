import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.Gateway("example", {
    gatewayIpAddress: "1.2.3.4",
    gatewayName: "example",
    gatewayTimezone: "GMT",
    gatewayType: "STORED",
});

