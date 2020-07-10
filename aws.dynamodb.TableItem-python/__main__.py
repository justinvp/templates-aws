import pulumi
import pulumi_aws as aws

example_table = aws.dynamodb.Table("exampleTable",
    attributes=[{
        "name": "exampleHashKey",
        "type": "S",
    }],
    hash_key="exampleHashKey",
    read_capacity=10,
    write_capacity=10)
example_table_item = aws.dynamodb.TableItem("exampleTableItem",
    hash_key=example_table.hash_key,
    item="""{
  "exampleHashKey": {"S": "something"},
  "one": {"N": "11111"},
  "two": {"N": "22222"},
  "three": {"N": "33333"},
  "four": {"N": "44444"}
}

""",
    table_name=example_table.name)

