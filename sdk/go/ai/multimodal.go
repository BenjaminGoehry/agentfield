package ai

import "strings"

type Image struct {
	URL      string `json:"url,omitempty"`
	Data     string `json:"data,omitempty"`
	MIMEType string `json:"mime_type,omitempty"`
}

type Audio struct {
	URL    string `json:"url,omitempty"`
	Data   string `json:"data,omitempty"`
	Format string `json:"format,omitempty"`
}

func detectMIMEType(path string) string {
	switch {
	case strings.HasSuffix(path, ".png"):
		return "image/png"
	case strings.HasSuffix(path, ".jpg"), strings.HasSuffix(path, ".jpeg"):
		return "image/jpeg"
	case strings.HasSuffix(path, ".gif"):
		return "image/gif"
	case strings.HasSuffix(path, ".webp"):
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}

func detectAudioFormat(path string) string {
	switch {
	case strings.HasSuffix(path, ".mp3"):
		return "mp3"
	case strings.HasSuffix(path, ".wav"):
		return "wav"
	case strings.HasSuffix(path, ".ogg"):
		return "ogg"
	case strings.HasSuffix(path, ".m4a"):
		return "m4a"
	case strings.HasSuffix(path, ".flac"):
		return "flac"
	default:
		return "unknown"
	}
}
