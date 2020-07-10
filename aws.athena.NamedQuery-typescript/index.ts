import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hogeBucket = new aws.s3.Bucket("hoge", {});
const testKey = new aws.kms.Key("test", {
    deletionWindowInDays: 7,
    description: "Athena KMS Key",
});
const testWorkgroup = new aws.athena.Workgroup("test", {
    configuration: {
        resultConfiguration: {
            encryptionConfiguration: {
                encryptionOption: "SSE_KMS",
                kmsKeyArn: testKey.arn,
            },
        },
    },
});
const hogeDatabase = new aws.athena.Database("hoge", {
    bucket: hogeBucket.id,
    name: "users",
});
const foo = new aws.athena.NamedQuery("foo", {
    database: hogeDatabase.name,
    query: pulumi.interpolate`SELECT * FROM ${hogeDatabase.name} limit 10;`,
    workgroup: testWorkgroup.id,
});

