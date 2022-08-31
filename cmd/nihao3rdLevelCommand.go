/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// nihao3rdLevelCommandCmd represents the nihao3rdLevelCommand command
var nihao3rdLevelCommandCmd = &cobra.Command{
	Use:   "nihao3rdLevelCommand",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	//参数验证器, 使用cobra提供的或者自己实现
	//Args: cobra.ExactArgs(2),
	Args: func(cmd *cobra.Command, args []string) error {
		nArgsRequired := 2
		if len(args) != nArgsRequired {
			return errors.New(fmt.Sprintf("requires exact %d args; given %v", nArgsRequired, args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("nihao3rdLevelCommand called. args are %v", args)

		first, _ := strconv.Atoi(args[0])
		second, _ := strconv.Atoi(args[1])

		var result int

		operationString, _ := cmd.Flags().GetString("operation")
		switch operationString {
		case "+":
			result = first + second
		case "-":
			result = first - second
		case "*":
			result = first * second
		case "/":
			result = first / second
		}
		t, _ := cmd.Flags().GetBool("verbose")
		if t {
			fmt.Printf("%d %s %d = ", first, operationString, second)
		}
		fmt.Println(result)
	},
}

func init() {
	helloSubcommandCmd.AddCommand(nihao3rdLevelCommandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nihao3rdLevelCommandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nihao3rdLevelCommandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	nihao3rdLevelCommandCmd.Flags().BoolP("verbose", "v", false, "print detailed info")
	nihao3rdLevelCommandCmd.Flags().StringP("operation", "o", "", "which operation")
	err := nihao3rdLevelCommandCmd.MarkFlagRequired("operation")
	if err != nil {
		log.Fatal("Set flag required failed")
	}
}
