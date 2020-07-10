import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sfnActivity = new aws.sfn.Activity("sfn_activity", {});

