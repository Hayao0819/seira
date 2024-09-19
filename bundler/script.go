package bundler

import (
	"io"
	"log/slog"
	"os"
	"path"

	"github.com/Hayao0819/seira/script"
	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
)

func resolveSources(i *script.Script, resolved *[]string) ([]*script.Script, error) {
	rt := []*script.Script{}
	if resolved == nil {
		resolved = &[]string{}
	}
	slog.Info("resolved", "file", *resolved)

	base := path.Dir(i.File.Name)
	for _, ipt := range *i.Imports {
		iptAbs := ""
		if path.IsAbs(ipt) {
			iptAbs = ipt
		} else {
			iptAbs = path.Join(base, ipt)
		}

		if lo.Contains(*resolved, iptAbs) {
			continue
		}
		slog.Info("resolving", "file", iptAbs)

		srcFile, err := os.Open(path.Join(base, ipt))
		if err != nil {
			if os.IsNotExist(err) {
				slog.Warn("file not found", "file", iptAbs)
				continue
			} else {
				return nil, err
			}
		}
		defer srcFile.Close()

		*resolved = append(*resolved, iptAbs)

		info, err := script.GetInfo(srcFile, path.Join(base, ipt))
		if err != nil {
			return nil, err
		}

		recursion, err := resolveSources(info, resolved)
		if err != nil {
			return nil, err
		}

		rt = append(rt, info)
		rt = append(rt, recursion...)

	}

	return rt, nil
}

func getTargetFileList(input io.Reader, iname string) (*[]string, error) {
	info, err := script.GetInfo(input, iname)
	if err != nil {
		return nil, err
	}

	if !hasMainFunction(info) {
		return nil, errors.New("main function not found")
	}

	resolved, err := resolveSources(info, nil)
	if err != nil {
		return nil, err
	}

	rt := lo.Map(resolved, func(i *script.Script, index int) string {
		return i.FullPath
	})

	return &rt, nil
}
