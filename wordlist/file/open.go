package file

import (
	"context"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func Open(ctx context.Context, ext string) string {
	filter := []runtime.FileFilter{
		{
			DisplayName: ext + " files (*." + ext + ")",
			Pattern:     "*." + ext,
		},
	}

	dialogOptions := runtime.OpenDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "wordlist." + ext,
		Title:                      "保存",
		Filters:                    filter,
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		ResolvesAliases:            true,
		TreatPackagesAsDirectories: false,
	}

	selectedFiles, err := runtime.OpenMultipleFilesDialog(ctx, dialogOptions)
	if len(selectedFiles) == 0 {
		return ""
	}
	if err != nil {
		fmt.Println(err)
		return ""
	}

	selectedFile := selectedFiles[0]
	bytes, err := os.ReadFile(selectedFile)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
