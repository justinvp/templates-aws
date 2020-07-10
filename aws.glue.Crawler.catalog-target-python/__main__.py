import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    catalog_targets=[{
        "database_name": aws_glue_catalog_database["example"]["name"],
        "tables": [aws_glue_catalog_table["example"]["name"]],
    }],
    configuration="""{
  "Version":1.0,
  "Grouping": {
    "TableGroupingPolicy": "CombineCompatibleSchemas"
  }
}

""",
    database_name=aws_glue_catalog_database["example"]["name"],
    role=aws_iam_role["example"]["arn"],
    schema_change_policy={
        "deleteBehavior": "LOG",
    })

