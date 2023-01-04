package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bogem/id3v2/v2"
)

func main() {
	var dir string

	// flags declaration using flag package
	flag.StringVar(&dir, "dir", ".", "Specify dir to pass dir with mp3 files default is .")
	flag.Parse()
	folderToAlbum(dir)

}

func folderToAlbum(dir string) {
	fmt.Println("folderToAlbum(" + dir + ")")
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			folderToAlbum(dir + "/" + file.Name())
		}
		if strings.HasSuffix(file.Name(), ".mp3") {
			v := strings.Split(dir, "/")
			album := dir
			if len(v) > 0 {
				album = v[len(v)-1]
			}
			err := setAlbum(dir+"/"+file.Name(), album)
			if err != nil {
				fmt.Println(err.Error() + " " + dir + "/" + file.Name())
			}
		}
	}
}

func setAlbum(filepath, album string) error {
	fmt.Println(album + " " + filepath)
	tag, err := id3v2.Open(filepath, id3v2.Options{Parse: true})
	if err != nil {
		return fmt.Errorf("error while opening mp3 file: \n%s", err.Error())
	}
	defer tag.Close()

	// Set tags
	tag.SetAlbum(album)

	// Write tag to file.mp3
	if err = tag.Save(); err != nil {
		return fmt.Errorf("error while saving a tag: %s", err.Error())
	}
	return nil
}
