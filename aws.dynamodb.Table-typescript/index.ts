import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const basic_dynamodb_table = new aws.dynamodb.Table("basic-dynamodb-table", {
    attributes: [
        {
            name: "UserId",
            type: "S",
        },
        {
            name: "GameTitle",
            type: "S",
        },
        {
            name: "TopScore",
            type: "N",
        },
    ],
    billingMode: "PROVISIONED",
    globalSecondaryIndexes: [{
        hashKey: "GameTitle",
        name: "GameTitleIndex",
        nonKeyAttributes: ["UserId"],
        projectionType: "INCLUDE",
        rangeKey: "TopScore",
        readCapacity: 10,
        writeCapacity: 10,
    }],
    hashKey: "UserId",
    rangeKey: "GameTitle",
    readCapacity: 20,
    tags: {
        Environment: "production",
        Name: "dynamodb-table-1",
    },
    ttl: {
        attributeName: "TimeToExist",
        enabled: false,
    },
    writeCapacity: 20,
});

