import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucket = new aws.s3.Bucket("exampleBucket", {});
const exampleAccessPoint = new aws.s3.AccessPoint("exampleAccessPoint", {bucket: exampleBucket.id});

