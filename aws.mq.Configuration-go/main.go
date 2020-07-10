package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/mq"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mq.NewConfiguration(ctx, "example", &mq.ConfigurationArgs{
			Data:          pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v", "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n", "<broker xmlns=\"http://activemq.apache.org/schema/core\">\n", "  <plugins>\n", "    <forcePersistencyModeBrokerPlugin persistenceFlag=\"true\"/>\n", "    <statisticsBrokerPlugin/>\n", "    <timeStampingBrokerPlugin ttlCeiling=\"86400000\" zeroExpirationOverride=\"86400000\"/>\n", "  </plugins>\n", "</broker>\n", "\n")),
			Description:   pulumi.String("Example Configuration"),
			EngineType:    pulumi.String("ActiveMQ"),
			EngineVersion: pulumi.String("5.15.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

