package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/guardduty"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		master, err := guardduty.NewDetector(ctx, "master", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		myThreatIntelSetBucketObject, err := s3.NewBucketObject(ctx, "myThreatIntelSetBucketObject", &s3.BucketObjectArgs{
			Acl:     pulumi.String("public-read"),
			Bucket:  bucket.ID(),
			Content: pulumi.String(fmt.Sprintf("%v%v", "10.0.0.0/8\n", "\n")),
			Key:     pulumi.String("MyThreatIntelSet"),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewThreatIntelSet(ctx, "myThreatIntelSetThreatIntelSet", &guardduty.ThreatIntelSetArgs{
			Activate:   pulumi.Bool(true),
			DetectorId: master.ID(),
			Format:     pulumi.String("TXT"),
			Location: pulumi.All(myThreatIntelSetBucketObject.Bucket, myThreatIntelSetBucketObject.Key).ApplyT(func(_args []interface{}) (string, error) {
				bucket := _args[0].(string)
				key := _args[1].(string)
				return fmt.Sprintf("%v%v%v%v", "https://s3.amazonaws.com/", bucket, "/", key), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

