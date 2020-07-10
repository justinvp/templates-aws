import pulumi
import pulumi_aws as aws

default_cluster = aws.neptune.Cluster("defaultCluster",
    apply_immediately="true",
    backup_retention_period=5,
    cluster_identifier="neptune-cluster-demo",
    engine="neptune",
    iam_database_authentication_enabled="true",
    preferred_backup_window="07:00-09:00",
    skip_final_snapshot=True)
example = aws.neptune.ClusterInstance("example",
    apply_immediately="true",
    cluster_identifier=default_cluster.id,
    engine="neptune",
    instance_class="db.r4.large")
default_topic = aws.sns.Topic("defaultTopic")
default_event_subscription = aws.neptune.EventSubscription("defaultEventSubscription",
    event_categories=[
        "maintenance",
        "availability",
        "creation",
        "backup",
        "restoration",
        "recovery",
        "deletion",
        "failover",
        "failure",
        "notification",
        "configuration change",
        "read replica",
    ],
    sns_topic_arn=default_topic.arn,
    source_ids=[example.id],
    source_type="db-instance",
    tags={
        "env": "test",
    })

