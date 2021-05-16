// +build static

package httpserver

import (
	"embed"
	"io/fs"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

//go:embed dist
var assets embed.FS

func (s *Server) setAssets() {
	s.Handle(`/`, http.FileServer(
		&assetfs.AssetFS{
			Asset: assets.ReadFile,
			AssetDir: func(path string) (names []string, err error) {
				entries, err := assets.ReadDir(path)
				if err != nil {
					return nil, err
				}
				for _, e := range entries {
					names = append(names, e.Name())
				}
				return names, nil
			},
			AssetInfo: func(path string) (fs.FileInfo, error) {
				file, err := assets.Open(path)
				if err != nil {
					return nil, err
				}
				return file.Stat()
			},
			Prefix:   "dist",
			Fallback: "index.html",
		},
	))
}
