package commands

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/jozsefsallai/fastbin-cli/config"
	"github.com/jozsefsallai/fastbin-cli/utils"
	"github.com/urfave/cli"
)

func printUrls(key string) {
	conf := config.GetConfig()

	documentURL := conf.Server + "/" + key
	rawURL := conf.Server + "/raw/" + key

	clipboard.WriteAll(documentURL)

	fmt.Println("Snippet uploaded successfully!")
	fmt.Println("URL:", documentURL)
	fmt.Println("Raw:", rawURL)
}

// CreateSnippet is the function that creates a snippet on the
// remote server either from a file or from another command's
// output
func CreateSnippet(ctx *cli.Context) error {
	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		if ctx.NArg() != 1 {
			fmt.Println("Please provide a file to upload.")
			return nil
		}

		fileName := ctx.Args().Get(0)
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fmt.Println("The specified file does not exist.")
			return nil
		}

		data, err := ioutil.ReadFile(ctx.Args().Get(0))
		if err != nil {
			panic(err)
		}

		result, err := utils.Upload(string(data))
		if err != nil {
			log.Fatal(err)
		}

		printUrls(result)

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

	result, err := utils.Upload(string(output))
	if err != nil {
		log.Fatal(err)
	}

	printUrls(result)

	return nil
}
