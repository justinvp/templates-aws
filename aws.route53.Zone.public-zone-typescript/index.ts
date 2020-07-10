import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.route53.Zone("primary", {});

