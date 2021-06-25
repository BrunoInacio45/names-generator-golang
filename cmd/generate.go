/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"names-generator/generate"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Use this project to generate aleatory names",
	Long: `Use this project to generate aleatory names. For example:

To generate 4 names with 2 surnames each: names-generator -n 4 -s 2;
To generate names in English: names-generator -l en;

Obs:
  To create names and surnames in anothers languages, it is necessary create 
the files inside 'data' folder.
`,
	Run: func(cmd *cobra.Command, args []string) {

		numberNames, _ := cmd.Flags().GetInt("number-names")
		numberSurnames, _ := cmd.Flags().GetInt("number-surnames")
		language, _ := cmd.Flags().GetString("language")
		generate.Generate(numberNames, numberSurnames, language)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("number-names", "n", 1, "Number of names to generated")
	generateCmd.Flags().IntP("number-surnames", "s", 3, "Numbers of surname on each name")
	generateCmd.Flags().StringP("language", "l", "pt", "Language of names and surnames")
}
