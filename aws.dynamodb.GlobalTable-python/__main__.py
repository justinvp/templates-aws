import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

us_east_1 = pulumi.providers.Aws("us-east-1", region="us-east-1")
us_west_2 = pulumi.providers.Aws("us-west-2", region="us-west-2")
us_east_1_table = aws.dynamodb.Table("us-east-1Table",
    attributes=[{
        "name": "myAttribute",
        "type": "S",
    }],
    hash_key="myAttribute",
    read_capacity=1,
    stream_enabled=True,
    stream_view_type="NEW_AND_OLD_IMAGES",
    write_capacity=1,
    opts=ResourceOptions(provider="aws.us-east-1"))
us_west_2_table = aws.dynamodb.Table("us-west-2Table",
    attributes=[{
        "name": "myAttribute",
        "type": "S",
    }],
    hash_key="myAttribute",
    read_capacity=1,
    stream_enabled=True,
    stream_view_type="NEW_AND_OLD_IMAGES",
    write_capacity=1,
    opts=ResourceOptions(provider="aws.us-west-2"))
my_table = aws.dynamodb.GlobalTable("myTable", replicas=[
    {
        "regionName": "us-east-1",
    },
    {
        "regionName": "us-west-2",
    },
],
opts=ResourceOptions(provider="aws.us-east-1",
    depends_on=[
        "aws_dynamodb_table.us-east-1",
        "aws_dynamodb_table.us-west-2",
    ]))

