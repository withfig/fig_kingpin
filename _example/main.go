package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	// 1. Add the figkingpin package
	figkingpin "github.com/withfig/fig_kingpin"
)

var (
	app = kingpin.New("Bucket Uploader", "A command-line app to upload files to a cloud storage bucket.")

	// 2. Add a top level flag to gen fig spec, it is hidden from the help output
	completionSpecFig = app.Flag("completion-spec-fig", "Generate completion script for fig.").Hidden().PreAction(figkingpin.GenerateFigCompletionSpec(app)).String()

	// Other commands
	uploadCmd  = app.Command("upload", "Upload a file to a specified bucket.")
	filePath   = uploadCmd.Arg("file", "Path of the file to upload.").Required().String()
	bucketName = uploadCmd.Flag("bucket", "Name of the bucket where the file will be uploaded.").Required().Enum("bucket1", "bucket2")

	nestedLevel1Cmd  = app.Command("nested", "A nested command.")
	nestedLevel2Cmd  = nestedLevel1Cmd.Command("level2", "A nested command.")
	nestedLevel2Flag = nestedLevel2Cmd.Flag("nested-flag", "A nested flag.").String()
	nestedLevel3Cmd  = nestedLevel2Cmd.Command("level3", "A nested command.")
)

func main() {
	var command = kingpin.MustParse(app.Parse(os.Args[1:]))

	switch command {
	case uploadCmd.FullCommand():
		//
	}
}
