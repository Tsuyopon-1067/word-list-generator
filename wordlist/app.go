package main

import (
	"context"
	"fmt"
	"wordlist/file"
	"wordlist/generator"
	"wordlist/parse"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GenerateTable(csv string) string {
	words, err := parse.ParseCsv(csv)
	if err != nil {
		return err.Error()
	}
	return generator.GenerateTableHtml(words)
}

func (a *App) GenerateCsv(html string) string {
	words, err := parse.ParseHtml(html)
	if err != nil {
		return err.Error()
	}
	return generator.GenerateCsvText(words)
}

func (a *App) Save(content string, ext string) {
	file.Save(a.ctx, content, ext)
}

func (a *App) Open(ext string) string {
	return file.Open(a.ctx, ext)
}
