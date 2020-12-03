package commands

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/jozsefsallai/fastbin-cli/config"
	"github.com/jozsefsallai/fastbin-cli/utils"
	"github.com/urfave/cli"
)

func printUrlsFunc(key string, mode string, extension string) {
	conf := config.GetConfigFunc()

	documentURL := fmt.Sprintf("%s/%s%s", conf.Server, key, extension)
	rawURL := fmt.Sprintf("%s/raw/%s", conf.Server, key)

	switch mode {
	case "full":
		clipboard.WriteAll(documentURL)
		fmt.Println(documentURL)
	case "raw":
		clipboard.WriteAll(rawURL)
		fmt.Println(rawURL)
	default:
		clipboard.WriteAll(documentURL)

		fmt.Println("Snippet uploaded successfully!")
		fmt.Println("URL:", documentURL)
		fmt.Println("Raw:", rawURL)
	}
}

// CreateSnippet is the function that creates a snippet on the
// remote server either from a file or from another command's
// output
func CreateSnippetFunc(ctx *cli.Context) error {
	mode := ""
	isFull := ctx.Bool("full")
	isRaw := ctx.Bool("raw")

	if isFull && isRaw {
		fmt.Println("Please use either --full or --raw, not both.")
		return nil
	}

	if isFull {
		mode = "full"
	}

	if isRaw {
		mode = "raw"
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		if ctx.NArg() == 0 {
			fmt.Println("Please provide a file to upload.")
			return nil
		}

		fileName := ctx.Args().Get(0)
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fmt.Println("The specified file does not exist.")
			return nil
		}

		extension := filepath.Ext(fileName)

		data, err := ioutil.ReadFile(ctx.Args().Get(0))
		if err != nil {
			panic(err)
		}

		result, err := utils.UploadFunc(string(data))
		if err != nil {
			log.Fatal(err)
		}

		printUrlsFunc(result, mode, extension)

		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	result, err := utils.UploadFunc(string(output))
	if err != nil {
		log.Fatal(err)
	}

	printUrlsFunc(result, mode, "")

	return nil
}
