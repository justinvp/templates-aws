import pulumi
import pulumi_aws as aws

default_instance = aws.rds.Instance("defaultInstance",
    allocated_storage=10,
    db_subnet_group_name="my_database_subnet_group",
    engine="mysql",
    engine_version="5.6.17",
    instance_class="db.t2.micro",
    name="mydb",
    parameter_group_name="default.mysql5.6",
    password="bar",
    username="foo")
default_topic = aws.sns.Topic("defaultTopic")
default_event_subscription = aws.rds.EventSubscription("defaultEventSubscription",
    event_categories=[
        "availability",
        "deletion",
        "failover",
        "failure",
        "low storage",
        "maintenance",
        "notification",
        "read replica",
        "recovery",
        "restoration",
    ],
    sns_topic=default_topic.arn,
    source_ids=[default_instance.id],
    source_type="db-instance")

