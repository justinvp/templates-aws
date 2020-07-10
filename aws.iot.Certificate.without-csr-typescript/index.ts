import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cert = new aws.iot.Certificate("cert", {
    active: true,
});

