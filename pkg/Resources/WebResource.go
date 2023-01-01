package Resources

import (
	"log"
	"os"
	"strings"
	"webgo/pkg/Mime"
)

/* A WebResource is either a HTML file, an image, sound, etc...	*/

type WebResource struct {
	LocalPath   string
	ContentType string
	loaded      bool
	content     string
}

func (wr *WebResource) Load() {
	data, err := os.ReadFile(wr.LocalPath)

	if err != nil {
		log.Fatalf("Could not read %v file (%v) => %v", wr, wr.LocalPath, err)
	}

	aux := strings.Split(wr.LocalPath, ".")
	ext := aux[len(aux)-1]

	wr.ContentType = Mime.GetMime(ext)
	wr.content = string(data)
	wr.loaded = true
}

func (wr WebResource) GetResult() string {
	return wr.content
}
