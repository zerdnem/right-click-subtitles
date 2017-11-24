package lib

import (
	"log"

	"github.com/oz/osdb"
)

//Auth ...
func Auth() (*osdb.Client, error) {
	c, err := osdb.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	if err = c.LogIn("", "", ""); err != nil {
		log.Fatal(err)
	}

	return c, err

}

//DownloadSubtitle ...
func DownloadSubtitle(subs osdb.Subtitles) error {
	c, err := Auth()
	if err != nil {
		panic(err)
	}

	return c.Download(&subs[0])
}

//SearchSubtitlesInPath ...
func SearchSubtitlesInPath(path string, language []string) {
	c, err := Auth()
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.FileSearch(path, language)
	if err != nil {
		log.Fatal(err)
	}

	err = DownloadSubtitle(res)
	if err != nil {
		panic(err)
	}
}
