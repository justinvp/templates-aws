import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appsync.GraphQLApi("example", {
    authenticationType: "AWS_IAM",
    schema: `schema {
	query: Query
}
type Query {
  test: Int
}
`,
});

