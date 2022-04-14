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

var regexpIPv4 = regexp.MustCompile(`((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])`)
var regexpIPv6 = regexp.MustCompile(`(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|[fF][eE]80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::([fF]{4}(:0{1,4}){0,1}:){0,1}
((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}
(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`)

var parseToggleIPv4 bool
var parseToggleIPv6 bool
var parseToggleStats bool

var parsedLines = 0
var foundAddresses = 0

func parseIPv4(cmd *cobra.Command, args []string) {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        ips := regexpIPv4.FindAllString(scanner.Text(), -1)
        for _, ip := range ips {
            fmt.Println(ip)
            foundAddresses += 1
        }
        parsedLines += 1
    }
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
    if parseToggleStats { parsePrintStats() }
}

func parsePrintStats() {
    fmt.Println("Lines parsed:", parsedLines)
    fmt.Println("Addresses found:", foundAddresses)
    fmt.Println(foundAddresses, "/", parsedLines)
}

func parseIPv6(cmd *cobra.Command, args []string) {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        ips := regexpIPv6.FindAllString(scanner.Text(), -1)
        for _, ip := range ips {
            fmt.Println(ip)
            foundAddresses += 1
        }
        parsedLines += 1
    }
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
    if parseToggleStats { parsePrintStats() }
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
    parseCmd.PersistentFlags().BoolVarP(
        &parseToggleIPv4, "ipv4", "4", true, "Parse IPv4 addresses",
    )
    parseCmd.PersistentFlags().BoolVarP(
        &parseToggleIPv6, "ipv6", "6", false, "Parse IPv6 addresses",
    )
    parseCmd.PersistentFlags().BoolVarP(
        &parseToggleStats, "stats", "s", false, "Provide parsed input stats.",
    )

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
