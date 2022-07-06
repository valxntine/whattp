package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/valxntine/whattp/internal/codes"
)

var border = "##########################################################################################"
var blankBorder = "#                                                                                        #"
var rootCmd = &cobra.Command{
	Use:   "whattp <status_code>",
	Short: "whattp is a CLI tool to find out what that pesky response code means.",
	Long:  `A minimal, offline-ready http response code explorer built with love by valxntine.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, ok := codes.StatusCodes[args[0]]
		if !ok {
			fmt.Fprintln(os.Stdout, center(border, 50))
			fmt.Fprintln(os.Stdout, center(blankBorder, 50))
			fmt.Fprintln(os.Stdout, center(addBorder(" _____________________________ "), 50))
			fmt.Fprintln(os.Stdout, center(addBorder("< are you sure thats...right? >"), 50))
			fmt.Fprintln(os.Stdout, center(addBorder(" ----------------------------- "), 50))
			fmt.Fprintln(os.Stdout, center(addBorder("\\   ^__^"), 50))
			fmt.Fprintln(os.Stdout, center(addBorder("         \\  (oo)\\_______"), 50))
			fmt.Fprintln(os.Stdout, center(addBorder("                (__)\\       )\\/\\"), 50))
			fmt.Fprintln(os.Stdout, center(addBorder("                ||----w |"), 50))
			fmt.Fprintln(os.Stdout, center(addBorder("                ||     ||"), 50))
			fmt.Fprintln(os.Stdout, center(blankBorder, 50))
			fmt.Fprintln(os.Stdout, center(border, 50))
		} else {
			fmt.Fprintln(os.Stdout, center(border, 50))
			fmt.Fprintln(os.Stdout, center(blankBorder, 50))
			fmt.Fprintln(os.Stdout, center(addBorder(v.LineOne), 50))
			fmt.Fprintln(os.Stdout, center(blankBorder, 50))
			fmt.Fprintln(os.Stdout, center(addBorder(v.LineTwo), 50))
			fmt.Fprintln(os.Stdout, center(addBorder(v.LineThree), 50))
			fmt.Fprintln(os.Stdout, center(addBorder(v.LineFour), 50))
			fmt.Fprintln(os.Stdout, center(addBorder(v.LineFive), 50))
			fmt.Fprintln(os.Stdout, center(addBorder(v.LineSix), 50))
			fmt.Fprintln(os.Stdout, center(addBorder(v.LineSeven), 50))
			fmt.Fprintln(os.Stdout, center(addBorder(v.LineEight), 50))
			fmt.Fprintln(os.Stdout, center(blankBorder, 50))
			fmt.Fprintln(os.Stdout, center(border, 50))

		}
	},
}

func center(s string, w int) string {
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}

func calculateGaps(s string) (int, int) {
	if len(s)%2 == 0 {
		return ((90 - len(s)) / 2) - 1, ((90 - len(s)) / 2) - 1
	}
	return ((89 - len(s)) / 2) - 1, ((89 - len(s)) / 2)
}

func addBorder(s string) string {
	leftGap, rightGap := calculateGaps(s)
	borderChar := "#"
	return fmt.Sprint(borderChar + strings.Repeat(" ", leftGap) + s + strings.Repeat(" ", rightGap) + borderChar)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
