/*
Copyright © 2022 x123 x123@users.noreply.github.com

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

    "github.com/elmasy-com/randomip"
	"github.com/spf13/cobra"
)

var generateFlagCount int
var generateToggleIPv4 bool
var generateToggleIPv6 bool

func generateIPv4(cmd *cobra.Command, args []string) {
    for i := 0; i < generateFlagCount; i++ {
        fmt.Println(randomip.GetPublicIPv4())
    }
}

func generateIPv6(cmd *cobra.Command, args []string) {
    for i := 0; i < generateFlagCount; i++ {
        fmt.Println(randomip.GetPublicIPv6())
    }
}

func generateGeneral(cmd *cobra.Command, args []string) {
    if generateToggleIPv4 { generateIPv4(cmd, args) }
    if generateToggleIPv6 { generateIPv6(cmd, args) }
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate random addresses",
	Long: `generate random addresses for use in testing`,
	Run: generateGeneral,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")
	generateCmd.PersistentFlags().IntVarP(
        &generateFlagCount, "count", "c", 1,
        "generate count (i.e., generate count items)")

    generateCmd.PersistentFlags().BoolVarP(&generateToggleIPv4, "ipv4", "4", true, "Generate IPv4 addresses")
    generateCmd.PersistentFlags().BoolVarP(&generateToggleIPv6, "ipv6", "6", false, "Generate IPv6 addresses")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}