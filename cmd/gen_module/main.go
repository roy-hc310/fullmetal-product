package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	snake := flag.String("snake", "", "feature name in snake_case, e.g. product_variant")
	camel := flag.String("camel", "", "feature name in CamelCase, e.g. ProductVariant")
	flag.Parse()

	if *snake == "" || *camel == "" {
		fmt.Println("Usage: go run main.go -snake product_variant -camel ProductVariant")
		os.Exit(1)
	}

	templateRoot := "pkg/gen/module"
	outputRoot := filepath.Join("modules", "module_"+*snake)

	err := filepath.WalkDir(templateRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		funcs := template.FuncMap{
			"Snake": func() string { return *snake },
			"Camel": func() string { return *camel },
		}

		tpl, err := template.New(filepath.Base(path)).Funcs(funcs).Parse(string(content))
		if err != nil {
			return err
		}

		relativePath := strings.TrimPrefix(path, templateRoot)
		outputPath := filepath.Join(outputRoot, relativePath)

		if strings.HasSuffix(outputPath, ".tpl") {
			outputPath = strings.TrimSuffix(outputPath, ".tpl")
		}

		if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		err = tpl.Execute(outFile, map[string]string{
			"Snake": *snake,
			"Camel": *camel,
		})
		if err != nil {
			return err
		}

		fmt.Println("✅ Generated:", outputPath)
		return nil
	})

	if err != nil {
		panic(err)
	}
}
