import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTable = new aws.dynamodb.Table("example", {
    attributes: [{
        name: "UserId",
        type: "S",
    }],
    hashKey: "UserId",
    readCapacity: 1,
    writeCapacity: 1,
});
const exampleRole = new aws.iam.Role("example", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "appsync.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
`,
});
const exampleRolePolicy = new aws.iam.RolePolicy("example", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "dynamodb:*"
      ],
      "Effect": "Allow",
      "Resource": [
        "${exampleTable.arn}"
      ]
    }
  ]
}
`,
    role: exampleRole.id,
});
const exampleGraphQLApi = new aws.appsync.GraphQLApi("example", {
    authenticationType: "API_KEY",
});
const exampleDataSource = new aws.appsync.DataSource("example", {
    apiId: exampleGraphQLApi.id,
    dynamodbConfig: {
        tableName: exampleTable.name,
    },
    serviceRoleArn: exampleRole.arn,
    type: "AMAZON_DYNAMODB",
});

