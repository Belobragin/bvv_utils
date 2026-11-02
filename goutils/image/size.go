package image

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	kb int64 = 1 << 10
	mb int64 = 1 << 20
	gb int64 = 1 << 30
)

// parseSize converts a human-readable size such as "100Kb", "5Mb" or "1Gb" to bytes.
func parseSize(s string) (int64, error) {
	v := strings.ToLower(strings.TrimSpace(s))

	units := []struct {
		suffix string
		mult   int64
	}{{"gb", gb}, {"mb", mb}, {"kb", kb}, {"b", 1}}

	for _, u := range units {
		if !strings.HasSuffix(v, u.suffix) {
			continue
		}
		num, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(v, u.suffix)), 64)
		if err != nil || num < 0 {
			break
		}
		return int64(num * float64(u.mult)), nil
	}

	return 0, fmt.Errorf("invalid size %q, expected format like 100Kb, 10Mb or 1Gb", s)
}

// formatSize renders bytes as Kb, Mb or Gb rounded to an integer.
func formatSize(size int64) string {
	switch {
	case size >= gb:
		return fmt.Sprintf("%dGb", int64(math.Round(float64(size)/float64(gb))))
	case size >= mb:
		return fmt.Sprintf("%dMb", int64(math.Round(float64(size)/float64(mb))))
	default:
		n := int64(math.Round(float64(size) / float64(kb)))
		if n == 0 && size > 0 {
			n = 1
		}
		return fmt.Sprintf("%dKb", n)
	}
}

// checkFileSize returns the file size in bytes and an error if it exceeds maxSize.
func checkFileSize(filePath, maxSize string) (int64, error) {
	limit, err := parseSize(maxSize)
	if err != nil {
		return 0, err
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	if info.IsDir() {
		return 0, fmt.Errorf("%s is a directory", filePath)
	}
	if info.Size() > limit {
		return 0, fmt.Errorf("file %s is %s, which exceeds the maximum size %s", filePath, formatSize(info.Size()), maxSize)
	}

	return info.Size(), nil
}

// GetFileSize returns the file size as a string like "122Kb", "12Mb" or "1Gb".
// An error is returned if the file is larger than maxSize (e.g. "100Gb").
func GetFileSize(filePath, maxSize string) (string, error) {
	size, err := checkFileSize(filePath, maxSize)
	if err != nil {
		return "", err
	}

	return formatSize(size), nil
}

// GetFileDim returns the image width and height in pixels (png, jpeg and gif are supported).
// An error is returned if the file is larger than maxSize (e.g. "100Mb").
func GetFileDim(filePath, maxSize string) (int, int, error) {
	if _, err := checkFileSize(filePath, maxSize); err != nil {
		return 0, 0, err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return 0, 0, fmt.Errorf("cannot read image dimensions of %s: %w", filePath, err)
	}

	return cfg.Width, cfg.Height, nil
}
