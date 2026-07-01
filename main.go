package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	inputFile := "confluent_kafka_setup.mp4"

	// Check whether the input file exists.
	// os.Stat() returns information about the file.
	// If the file doesn't exist, print an error and stop the program.
	if _, err := os.Stat(inputFile); err != nil {
		fmt.Println("Input file not found:", err)
		return
	}

	os.MkdirAll("output", 0755)

	args := []string{
		// Input video file
		"-i", inputFile,
		// Encode video using the H.264 codec
		"-codec:v", "libx264",
		// Encode audio using the AAC codec
		"-codec:a", "aac",
		// Include all streams (video, audio, subtitles if present)
		"-map", "0",
		// Set the output format to HLS (HTTP Live Streaming)
		"-f", "hls",
		// Split the video into 10-second segments
		"-hls_time", "10",
		// Include all generated segments in the playlist
		// 0 means do not limit the playlist size.
		"-hls_list_size", "0",
		// Pattern used for naming segment files.
		// %03d creates names like:
		// segment_000.ts
		// segment_001.ts
		// segment_002.ts
		"-hls_segment_filename", "output/segment_%03d.ts",
		// Output playlist file.
		// This file references all the generated .ts segments.
		"output/playlist.m3u8",
	}
	// Create an FFmpeg command with the arguments above.
	// This is equivalent to running the command in a terminal:
	//
	// ffmpeg -i confluent_kafka_setup.mp4
	//        -codec:v libx264
	//        -codec:a aac
	//        -map 0
	//        -f hls
	//        -hls_time 10
	//        -hls_list_size 0
	//        -hls_segment_filename output/segment_%03d.ts
	//        output/playlist.m3u8
	cmd := exec.Command("ffmpeg", args...)

	// Execute the FFmpeg command.
	// CombinedOutput() captures both:
	// - Standard Output (stdout)
	// - Error Output (stderr)
	// This makes debugging easier if FFmpeg fails.
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("FFmpeg error:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println(string(output))
	fmt.Println("Conversion successful!")
}
