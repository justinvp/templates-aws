import pulumi
import pulumi_aws as aws

config = pulumi.Config()
broker_id = config.get("brokerId")
if broker_id is None:
    broker_id = ""
broker_name = config.get("brokerName")
if broker_name is None:
    broker_name = ""
by_id = aws.mq.get_broker(broker_id=broker_id)
by_name = aws.mq.get_broker(broker_name=broker_name)

