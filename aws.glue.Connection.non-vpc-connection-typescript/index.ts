import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Connection("example", {
    connectionProperties: {
        JDBC_CONNECTION_URL: "jdbc:mysql://example.com/exampledatabase",
        PASSWORD: "examplepassword",
        USERNAME: "exampleusername",
    },
});

