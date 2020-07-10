using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Aws.ElasticTranscoder.Preset("bar", new Aws.ElasticTranscoder.PresetArgs
        {
            Audio = new Aws.ElasticTranscoder.Inputs.PresetAudioArgs
            {
                AudioPackingMode = "SingleTrack",
                BitRate = "96",
                Channels = "2",
                Codec = "AAC",
                SampleRate = "44100",
            },
            AudioCodecOptions = new Aws.ElasticTranscoder.Inputs.PresetAudioCodecOptionsArgs
            {
                Profile = "AAC-LC",
            },
            Container = "mp4",
            Description = "Sample Preset",
            Thumbnails = new Aws.ElasticTranscoder.Inputs.PresetThumbnailsArgs
            {
                Format = "png",
                Interval = "120",
                MaxHeight = "auto",
                MaxWidth = "auto",
                PaddingPolicy = "Pad",
                SizingPolicy = "Fit",
            },
            Video = new Aws.ElasticTranscoder.Inputs.PresetVideoArgs
            {
                BitRate = "1600",
                Codec = "H.264",
                DisplayAspectRatio = "16:9",
                FixedGop = "false",
                FrameRate = "auto",
                KeyframesMaxDist = "240",
                MaxFrameRate = "60",
                MaxHeight = "auto",
                MaxWidth = "auto",
                PaddingPolicy = "Pad",
                SizingPolicy = "Fit",
            },
            VideoCodecOptions = 
            {
                { "ColorSpaceConversionMode", "None" },
                { "InterlacedMode", "Progressive" },
                { "Level", "2.2" },
                { "MaxReferenceFrames", "3" },
                { "Profile", "main" },
            },
            VideoWatermarks = 
            {
                new Aws.ElasticTranscoder.Inputs.PresetVideoWatermarkArgs
                {
                    HorizontalAlign = "Right",
                    HorizontalOffset = "10px",
                    Id = "Test",
                    MaxHeight = "20%",
                    MaxWidth = "20%",
                    Opacity = "55.5",
                    SizingPolicy = "ShrinkToFit",
                    Target = "Content",
                    VerticalAlign = "Bottom",
                    VerticalOffset = "10px",
                },
            },
        });
    }

}

