package cmd

func Execute() error {
	root := rootCmd()
	return root.Execute()
}
