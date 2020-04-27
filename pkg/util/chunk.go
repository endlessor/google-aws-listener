package util

import (
	"google-rtb/config"
	"google-rtb/pkg/svc/requestsCounter"
)

func IsChunkFull() bool {
	if requestsCounter.GetNumber() > 0 && (requestsCounter.GetNumber()%config.Cfg.FileUploader.ChunkSize == 0) {
		return true
	}

	return false
}
