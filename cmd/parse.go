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
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var regexpIPv4 = regexp.MustCompile(`((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])`)
var regexpIPv6 = regexp.MustCompile(`((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:)))(%.+)?(\/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9]))?`)

var parseToggleIPv4 bool
var parseToggleIPv6 bool
var parseToggleStats bool

var parsedLines = 0
var foundIPv4Addresses = 0
var foundIPv6Addresses = 0

func parseIPv4(cmd *cobra.Command, args []string) error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ips := regexpIPv4.FindAllString(scanner.Text(), -1)
		for _, ip := range ips {
			fmt.Println(ip)
			foundIPv4Addresses += 1
		}
		parsedLines += 1
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func parseIPv6(cmd *cobra.Command, args []string) error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ips := regexpIPv6.FindAllString(scanner.Text(), -1)
		for _, ip := range ips {
			fmt.Println(ip)
			foundIPv6Addresses += 1
		}
		parsedLines += 1
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func parsePrintStats() error {
	if parsedLines == 0 {
		return errors.New("no lines could be parsed from stdin")
	}
	fmt.Println("# Statistics")
	fmt.Println("# Lines parsed:", parsedLines)
	if parseToggleIPv4 {
		fmt.Printf("# IPv4 Addresses found/Lines parsed: %d/%d\n", foundIPv4Addresses, parsedLines)
		fmt.Printf("# Percentage of lines with IPv4 addresses: %.2f%%\n", float64(foundIPv4Addresses)/float64(parsedLines)*100.0)
		fmt.Printf("# Average number of IPv4 addresses found/Line: %.2f\n", float64(foundIPv4Addresses)/float64(parsedLines))
	}
	if parseToggleIPv6 {
		fmt.Printf("# IPv6 Addresses found/Lines parsed: %d/%d\n", foundIPv6Addresses, parsedLines)
		fmt.Printf("# Percentage of lines with IPv6 addresses: %.2f%%\n", float64(foundIPv6Addresses)/float64(parsedLines)*100.0)
		fmt.Printf("# Average number of IPv6 addresses found/Line: %.2f\n", float64(foundIPv6Addresses)/float64(parsedLines))
	}
	return nil
}

func parseGeneral(cmd *cobra.Command, args []string) {
	if parseToggleIPv4 {
		if err := parseIPv4(cmd, args); err != nil {
			fmt.Println(err)
			return
		}
	}
	if parseToggleIPv6 {
		if err := parseIPv6(cmd, args); err != nil {
			fmt.Println(err)
			return
		}
	}
	if parseToggleStats {
		if err := parsePrintStats(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

// parseCmd represents the ipv4 command
var parseCmd = &cobra.Command{
	Use: "parse",
	Args: func(cmd *cobra.Command, args []string) error {
		if parseToggleIPv4 && parseToggleIPv6 {
			return errors.New("only one address flag (--ipv4 or --ipv6) may be active")
		}
		return nil
	},
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
