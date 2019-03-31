//Copyright (C) 2019  Yurizal Susanto  <yursan9@pm.me>
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"trash-cli/pkg/rm"

	"github.com/spf13/cobra"
)

var (
	interactive bool
	pattern     string
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the file from the trash",
	Long: `Permanently remove files from trash. Enter former absolute path
of files or use pattern to delete files that match the pattern.`,
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rm.Run(args, pattern, interactive)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	rmCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Prompt before every removal")
	rmCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Remove file that match the pattern")
}
