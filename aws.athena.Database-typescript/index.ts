import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hogeBucket = new aws.s3.Bucket("hoge", {});
const hogeDatabase = new aws.athena.Database("hoge", {
    bucket: hogeBucket.bucket,
    name: "database_name",
});

