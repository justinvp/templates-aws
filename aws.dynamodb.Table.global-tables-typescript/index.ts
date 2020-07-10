import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dynamodb.Table("example", {
    attributes: [{
        name: "TestTableHashKey",
        type: "S",
    }],
    billingMode: "PAY_PER_REQUEST",
    hashKey: "TestTableHashKey",
    replicas: [
        {
            regionName: "us-east-2",
        },
        {
            regionName: "us-west-2",
        },
    ],
    streamEnabled: true,
    streamViewType: "NEW_AND_OLD_IMAGES",
});

