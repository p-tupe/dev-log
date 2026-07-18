//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"slices"
	"strings"
	"time"
)

func main() {
	output, err := exec.Command("git", "status", "--short").Output()
	if err != nil {
		log.Fatalln("Error reading notes dir: ", err)
	}

	for note := range strings.FieldsSeq(string(output)) {
		ext := path.Ext(note)
		if ext != ".md" && ext != ".mdx" {
			continue
		}

		lastModified, err := exec.Command("git", "log", "-1", "--pretty=format:%ad", "--date=default", note).Output()
		if err != nil {
			log.Fatalf("Error while fetching last modified for %s: %s\n", note, err)
		}

		if len(lastModified) == 0 {
			lastModified = []byte(time.Now().Local().String())
		}

		contents, err := os.ReadFile(note)
		if err != nil {
			log.Fatalf("Error while reading note %s: %s\n", note, err)
		}

		if slices.Equal(contents[:3], []byte("---")) {
			re := regexp.MustCompile("^---\nmodified:.*\n---\n")
			contents = re.ReplaceAll(contents, []byte(""))
		}

		newContent := fmt.Appendf([]byte(""), "---\nmodified: %s\n---\n%s", lastModified, contents)
		err = os.WriteFile(note, newContent, 0o644)
		if err != nil {
			log.Fatalf("Error while writing contents to note %s: %s\n", note, err)
		}
	}
}
