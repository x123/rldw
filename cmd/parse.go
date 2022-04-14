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
    "bufio"
    "fmt"
    "os"
    "regexp"

    "github.com/spf13/cobra"
)

var regexpIPV4 = regexp.MustCompile("[0-9]{1,3}[.][0-9]{1,3}[.][0-9]{1,3}[.][0-9]{1,3}")

var parseToggleIPv4 bool
var parseToggleIPv6 bool

func parseIPv4(cmd *cobra.Command, args []string) {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        ips := regexpIPV4.FindAllString(scanner.Text(), -1)
        for _, ip := range ips {
            fmt.Println(ip)
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func parseIPv6(cmd *cobra.Command, args []string) {
    fmt.Println("IPv6 parsing not yet supported")
}

func parseGeneral(cmd *cobra.Command, args []string) {
    if parseToggleIPv4 { parseIPv4(cmd, args) }
    if parseToggleIPv6 { parseIPv6(cmd, args) }
}

// parseCmd represents the ipv4 command
var parseCmd = &cobra.Command{
    Use:   "parse",
    Short: "Strip IP addesses from stdin",
    Long: `Strip IP addresses from stdin, for further parsing.

Example for IPv4 addressess:
$ echo "4.2.2.2 sometexthere 8.8.8.8" | rldw parse -4
4.2.2.2
8.8.8.8`,
    Run: parseGeneral,
}

func init() {
    rootCmd.AddCommand(parseCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // parseCmd.PersistentFlags().String("foo", "", "A help for foo")
    parseCmd.PersistentFlags().BoolVarP(&parseToggleIPv4, "ipv4", "4", true, "Parse IPv4 addresses")
    parseCmd.PersistentFlags().BoolVarP(&parseToggleIPv6, "ipv6", "6", false, "Parse IPv6 addresses")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
