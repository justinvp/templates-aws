import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    catalogTargets: [{
        databaseName: aws_glue_catalog_database_example.name,
        tables: [aws_glue_catalog_table_example.name],
    }],
    configuration: `{
  "Version":1.0,
  "Grouping": {
    "TableGroupingPolicy": "CombineCompatibleSchemas"
  }
}
`,
    databaseName: aws_glue_catalog_database_example.name,
    role: aws_iam_role_example.arn,
    schemaChangePolicy: {
        deleteBehavior: "LOG",
    },
});

