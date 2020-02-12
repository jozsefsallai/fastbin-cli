package commands

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/jozsefsallai/fastbin-cli/utils"
	"github.com/urfave/cli"
)

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
			panic(err)
		}

		fmt.Println(string(result))

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

	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}

	return nil
}
