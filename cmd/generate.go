//Package cmd generate
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
	if generateToggleIPv4 {
		generateIPv4(cmd, args)
	}
	if generateToggleIPv6 {
		generateIPv6(cmd, args)
	}
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate random addresses",
	Long: `generate random addresses for use in testing

Generate 2 random IPv4 addresses:
    $ rldw generate -4 -c 2
    7.26.13.106
    212.212.148.53

Generate 3 random IPv6 addresses:
    $ rldw generate -4=false -6 -c 3
    5dae:937a:4b94:8253:ef36:bdb6:4752:68df
    e92c:ad08:e2a2:6096:77e9:befe:1a02:ee5e
    ec96:140:ff0e:7ba7:6537:5af2:7801:155

Generate 1 IPv4 address *and* one IPv6 address:
    rldw generate -4 -6 -c1
    98.234.210.100
    5c47:f78c:59a4:fe43:dcc0:3085:1e1b:dadd
`,
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
