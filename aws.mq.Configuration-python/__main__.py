import pulumi
import pulumi_aws as aws

example = aws.mq.Configuration("example",
    data="""<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<broker xmlns="http://activemq.apache.org/schema/core">
  <plugins>
    <forcePersistencyModeBrokerPlugin persistenceFlag="true"/>
    <statisticsBrokerPlugin/>
    <timeStampingBrokerPlugin ttlCeiling="86400000" zeroExpirationOverride="86400000"/>
  </plugins>
</broker>

""",
    description="Example Configuration",
    engine_type="ActiveMQ",
    engine_version="5.15.0")

