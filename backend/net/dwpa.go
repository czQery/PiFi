package net

import (
	"context"
	"os"
)

func DWPAGet() []string {
	dir, dirErr := os.ReadDir("./cap")
	if dirErr != nil {
		return nil
	}

	var list []string

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil || info.Size() <= 128 { // skip empty files
			continue
		}

		list = append(list, entry.Name())
	}

	return list
}

func DWPAUpload(ctx context.Context) error {
	return nil
}
