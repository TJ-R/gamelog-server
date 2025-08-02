package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "cli",
		Short: "gamelog-cli is an interactive way to use gamelog to track games",
		Long: `gamelog-cli The use case for this right now is to use it to load IGDB
			   info into the db.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Do Nothing")
		},
	}

	var pullIGDB = &cobra.Command{
		Use: "pullIGDB",
		Short: "Command to pull information from relevant tables in IGDB",
		Long: `Pulls information from the IGDB using apicalypse and stores it in
		       our datebase to use.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Do Something")

			// Start the process of pulling all of the relevant IGDB data into database
			// here the function to do it will be in the internal package somewhere
		},
	}

	rootCmd.AddCommand(pullIGDB)

	rootCmd.Execute()
	
	// Execute()
}



// func Execute() {
// 	for {
// 		if err := rootCmd.Execute(); err != nil {
// 			fmt.Fprintf(os.Stderr, err.Error())
// 			os.Exit(1)
// 		}
// 	}
// }

