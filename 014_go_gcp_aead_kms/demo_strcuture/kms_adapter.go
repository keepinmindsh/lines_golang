package demo_strcuture

import "C"

type KMSAdapter interface {
	*KMSGcp
}

type KMSValueT struct {
	cmpId     string
	plaintext []byte
	aead      []byte
}

type KMSValueR struct {
	result []byte
}

func NewKMSHandler[T KMSAdapter](kmsType KmsConstant) T {
	switch kmsType {
	case KMS_GCP:
		return &KMSGcp{}
	case KMS_AWS:
		return nil
	}

	return nil
}
