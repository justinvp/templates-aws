import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const us_east_1 = new aws.Provider("us-east-1", {
    region: "us-east-1",
});
const us_west_2 = new aws.Provider("us-west-2", {
    region: "us-west-2",
});
const us_east_1Table = new aws.dynamodb.Table("us-east-1", {
    attributes: [{
        name: "myAttribute",
        type: "S",
    }],
    hashKey: "myAttribute",
    readCapacity: 1,
    streamEnabled: true,
    streamViewType: "NEW_AND_OLD_IMAGES",
    writeCapacity: 1,
}, { provider: us_east_1 });
const us_west_2Table = new aws.dynamodb.Table("us-west-2", {
    attributes: [{
        name: "myAttribute",
        type: "S",
    }],
    hashKey: "myAttribute",
    readCapacity: 1,
    streamEnabled: true,
    streamViewType: "NEW_AND_OLD_IMAGES",
    writeCapacity: 1,
}, { provider: us_west_2 });
const myTable = new aws.dynamodb.GlobalTable("myTable", {
    replicas: [
        {
            regionName: "us-east-1",
        },
        {
            regionName: "us-west-2",
        },
    ],
}, { provider: us_east_1, dependsOn: [us_east_1Table, us_west_2Table] });

