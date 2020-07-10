import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mykey = new aws.kms.Key("mykey", {
    deletionWindowInDays: 10,
    description: "This key is used to encrypt bucket objects",
});
const mybucket = new aws.s3.Bucket("mybucket", {
    serverSideEncryptionConfiguration: {
        rule: {
            applyServerSideEncryptionByDefault: {
                kmsMasterKeyId: mykey.arn,
                sseAlgorithm: "aws:kms",
            },
        },
    },
});

