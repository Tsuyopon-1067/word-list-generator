package file

import (
	"context"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func Save(ctx context.Context, content string, ext string) {
	filter := []runtime.FileFilter{
		{
			DisplayName: ext + " files (*." + ext + ")",
			Pattern:     "*." + ext,
		},
	}

	dialogOptions := runtime.SaveDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "wordlist." + ext,
		Title:                      "保存",
		Filters:                    filter,
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: false,
	}

	selectedFile, err := runtime.SaveFileDialog(ctx, dialogOptions)
	if selectedFile == "" {
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	os.WriteFile(selectedFile, []byte(content), 0644)
}
