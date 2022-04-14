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
    "bufio"
    "os"

    "github.com/spf13/cobra"
)

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

// ipv4Cmd represents the ipv4 command
var ipv4Cmd = &cobra.Command{
    Use:   "ipv4",
    Short: "Strip IPv4 addesses from stdin",
    Long: `Strip IPv4 addresses from stdin, for further parsing. For example:

$ echo "4.2.2.2 sometexthere 8.8.8.8" | rldw ipv4
4.2.2.2
8.8.8.8

`,
    Run: parseIPv4,
}

func init() {
    rootCmd.AddCommand(ipv4Cmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // ipv4Cmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // ipv4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
