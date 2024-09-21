package bundler

import (
	"log/slog"
	"os"
	"path"
	"strings"

	cp "github.com/otiai10/copy"
	"github.com/samber/lo"
)

func copySourceFilesToDir(files []string, base string, work string) error {
	var err error
	lo.ForEach(files, func(file string, i int) {
		reldir := strings.TrimPrefix(path.Dir(file), base)
		if reldir == "" {
			reldir = "."
		}
		slog.Debug("copying file", "file", file, "dir", reldir)

		destDir := path.Join(work, reldir)

		if err := os.MkdirAll(destDir, 0755); err != nil {
			slog.Error("failed to create directory", "dir", destDir, "error", err)
		}

		destFile := path.Join(destDir, path.Base(file))

		if err := cp.Copy(file, destFile); err != nil {
			slog.Error("failed to copy file", "file", file, "dest", destFile, "error", err)
		}
	})
	return err
}
