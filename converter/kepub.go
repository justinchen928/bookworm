package converter

import (
	"archive/zip"
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/pgaskin/kepubify/v4/kepub"
)

func ToKepub(epub_path string, output_path string) {
	converter := kepub.NewConverter()
	fi, err := zip.OpenReader(epub_path)
	if err != nil {
		log.Panicf("error: %v", err)
	}
	defer fi.Close()

	fo, err := os.CreateTemp(filepath.Dir(output_path), ".kepubify."+filepath.Base(output_path)+".*")
	if err != nil {
		log.Panicf("error: %v", err)
	}
	defer os.Remove(fo.Name())
	if err := converter.Convert(context.Background(), fo, fi); err != nil {
		log.Panicf("error: %v", err)
	}
	if err := fo.Sync(); err != nil {
		log.Panicf("error: %v", err)
	}

	if err := fo.Close(); err != nil {
		log.Panicf("error: %v", err)
	}

	if err := os.Rename(fo.Name(), output_path); err != nil {
		log.Panicf("error: %v", err)
	}
}
