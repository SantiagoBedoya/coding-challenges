package cmd

import (
	"bufio"
	"fmt"
	"go-wc/pkg/counter"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-wc",
	Short: "Golang wc CLI",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		countBytes, _ := cmd.Flags().GetBool("countBytes")
		countLines, _ := cmd.Flags().GetBool("countLines")
		countWords, _ := cmd.Flags().GetBool("countWords")
		countCharacters, _ := cmd.Flags().GetBool("countCharacters")
		// fmt.Println(countBytes, countLines, countWords, countCharacters)

		content := ""
		if len(args) == 0 {
			reader := bufio.NewReader(os.Stdin)
			data := make([]string, 0)
			for {
				b, _, err := reader.ReadLine()
				data = append(data, string(b))
				if err != nil {
					if err == io.EOF {
						break
					}
					log.Fatal(err)
				}
			}
			content = strings.Join(data, "\r\n")
		} else {
			filePath := args[0]
			data, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatal(err)
			}
			content = string(data)
		}
		response := counter.Count(countBytes, countLines, countWords, countCharacters, content)
		if len(args) != 0 {
			response += fmt.Sprintf("\t%s", args[0])
		}
		fmt.Println(response)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("countBytes", "c", false, "count bytes")
	rootCmd.Flags().BoolP("countLines", "l", false, "count lines")
	rootCmd.Flags().BoolP("countWords", "w", false, "count words")
	rootCmd.Flags().BoolP("countCharacters", "m", false, "count characters")
}
