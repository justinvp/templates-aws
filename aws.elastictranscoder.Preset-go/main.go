package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elastictranscoder"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elastictranscoder.NewPreset(ctx, "bar", &elastictranscoder.PresetArgs{
			Audio: &elastictranscoder.PresetAudioArgs{
				AudioPackingMode: pulumi.String("SingleTrack"),
				BitRate:          pulumi.String("96"),
				Channels:         pulumi.String("2"),
				Codec:            pulumi.String("AAC"),
				SampleRate:       pulumi.String("44100"),
			},
			AudioCodecOptions: &elastictranscoder.PresetAudioCodecOptionsArgs{
				Profile: pulumi.String("AAC-LC"),
			},
			Container:   pulumi.String("mp4"),
			Description: pulumi.String("Sample Preset"),
			Thumbnails: &elastictranscoder.PresetThumbnailsArgs{
				Format:        pulumi.String("png"),
				Interval:      pulumi.String("120"),
				MaxHeight:     pulumi.String("auto"),
				MaxWidth:      pulumi.String("auto"),
				PaddingPolicy: pulumi.String("Pad"),
				SizingPolicy:  pulumi.String("Fit"),
			},
			Video: &elastictranscoder.PresetVideoArgs{
				BitRate:            pulumi.String("1600"),
				Codec:              pulumi.String("H.264"),
				DisplayAspectRatio: pulumi.String("16:9"),
				FixedGop:           pulumi.String("false"),
				FrameRate:          pulumi.String("auto"),
				KeyframesMaxDist:   pulumi.String("240"),
				MaxFrameRate:       pulumi.String("60"),
				MaxHeight:          pulumi.String("auto"),
				MaxWidth:           pulumi.String("auto"),
				PaddingPolicy:      pulumi.String("Pad"),
				SizingPolicy:       pulumi.String("Fit"),
			},
			VideoCodecOptions: pulumi.StringMap{
				"ColorSpaceConversionMode": pulumi.String("None"),
				"InterlacedMode":           pulumi.String("Progressive"),
				"Level":                    pulumi.String("2.2"),
				"MaxReferenceFrames":       pulumi.String("3"),
				"Profile":                  pulumi.String("main"),
			},
			VideoWatermarks: elastictranscoder.PresetVideoWatermarkArray{
				&elastictranscoder.PresetVideoWatermarkArgs{
					HorizontalAlign:  pulumi.String("Right"),
					HorizontalOffset: pulumi.String("10px"),
					Id:               pulumi.String("Test"),
					MaxHeight:        pulumi.String(fmt.Sprintf("%v%v", "20", "%")),
					MaxWidth:         pulumi.String(fmt.Sprintf("%v%v", "20", "%")),
					Opacity:          pulumi.String("55.5"),
					SizingPolicy:     pulumi.String("ShrinkToFit"),
					Target:           pulumi.String("Content"),
					VerticalAlign:    pulumi.String("Bottom"),
					VerticalOffset:   pulumi.String("10px"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

