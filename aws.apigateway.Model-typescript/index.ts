import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDemoAPI = new aws.apigateway.RestApi("MyDemoAPI", {
    description: "This is my API for demonstration purposes",
});
const myDemoModel = new aws.apigateway.Model("MyDemoModel", {
    contentType: "application/json",
    description: "a JSON schema",
    restApi: myDemoAPI.id,
    schema: `{
  "type": "object"
}
`,
});

