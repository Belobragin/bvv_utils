package config

type PictureConfigI interface {
	GetMultypartSize() int64
}

type PictureConfig struct {
	MultypartSize int64 `conf:"env:MAX_MULTIPART_SIZE"`
}

func (m *PictureConfig) GetMultypartSize() int64 {
	return m.MultypartSize
}

func (m *PictureConfig) ValidatePictureConfig() error {
	return nil
}
