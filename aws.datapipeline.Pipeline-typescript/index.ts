import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultPipeline = new aws.datapipeline.Pipeline("default", {});

