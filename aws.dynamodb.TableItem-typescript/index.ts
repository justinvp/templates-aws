import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTable = new aws.dynamodb.Table("example", {
    attributes: [{
        name: "exampleHashKey",
        type: "S",
    }],
    hashKey: "exampleHashKey",
    readCapacity: 10,
    writeCapacity: 10,
});
const exampleTableItem = new aws.dynamodb.TableItem("example", {
    hashKey: exampleTable.hashKey,
    item: `{
  "exampleHashKey": {"S": "something"},
  "one": {"N": "11111"},
  "two": {"N": "22222"},
  "three": {"N": "33333"},
  "four": {"N": "44444"}
}
`,
    tableName: exampleTable.name,
});

