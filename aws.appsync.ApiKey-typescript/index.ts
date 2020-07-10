import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGraphQLApi = new aws.appsync.GraphQLApi("example", {
    authenticationType: "API_KEY",
});
const exampleApiKey = new aws.appsync.ApiKey("example", {
    apiId: exampleGraphQLApi.id,
    expires: "2018-05-03T04:00:00Z",
});

