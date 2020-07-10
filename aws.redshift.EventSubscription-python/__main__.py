import pulumi
import pulumi_aws as aws

default_cluster = aws.redshift.Cluster("defaultCluster",
    cluster_identifier="default",
    database_name="default")
default_topic = aws.sns.Topic("defaultTopic")
default_event_subscription = aws.redshift.EventSubscription("defaultEventSubscription",
    event_categories=[
        "configuration",
        "management",
        "monitoring",
        "security",
    ],
    severity="INFO",
    sns_topic_arn=default_topic.arn,
    source_ids=[default_cluster.id],
    source_type="cluster",
    tags={
        "Name": "default",
    })

