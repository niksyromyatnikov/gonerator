package generator

import (
	"fmt"
	"generator/blocks"
	"os"
	"os/exec"
)

type Generator struct {
	StorePath string
}

func NewGenerator() *Generator{
	return &Generator{"generated/"}
}

func (G *Generator) Generate(Listing blocks.CodeInterface) string{
	return Listing.Code()
}

func (G *Generator) GenerateFile(File *blocks.File) (string, error) {
	code := File.Code()

	fileName := File.GetFileName() + ".go"

	path := ""

	if path = File.GetFilePath(); path == "" {
		path = G.GetStoragePath()
	}

	err := os.MkdirAll(path, os.ModePerm)

	filePath := path + fileName

	f, err := os.Create(filePath)

	if err != nil {
		fmt.Print(err)
	}

	_, err = f.WriteString(code)

	if err != nil {
		fmt.Println(err)
	}

	if err == nil {
		err = format(filePath)
	}

	fmt.Println("Generated", filePath)

	return code, err
}

func (G *Generator) SetStoragePath(path string) *Generator {
	G.StorePath = path
	return G
}

func (G *Generator) GetStoragePath() string {
	return G.StorePath
}


func format(path string) error{
	cmd := exec.Command("gofmt", "-w", path)
	return cmd.Run()
}