package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

const (
	DefaultGolangUrl  = "https://golang.org"
	DefaultGorootsDir = "goroots"
	DefaultTempDir    = "tmp"
)

type Config struct {
	GolangUrl  string   `mapstructure:"golang_url"`
	GorootsDir string   `mapstructure:"goroots_dir"`
	TempDir    string   `mapstructure:"temp_dir"`
	Versions   []string `mapstructure:"versions"`
}

var cfg Config

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.2.0",
	Use:     "godl",
	Short:   "golang downloader",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is %s)", filepath.Join(defaultConfigPath(), "config.toml")))

	rootCmd.PersistentFlags().String("gourl", DefaultGorootsDir, "golang url")
	err := viper.BindPFlag("golang_url", rootCmd.PersistentFlags().Lookup("gourl"))
	cobra.CheckErr(err)

	rootCmd.PersistentFlags().String("goroots", DefaultGorootsDir, "goroots directory")
	err = viper.BindPFlag("goroots_dir", rootCmd.PersistentFlags().Lookup("goroots"))
	cobra.CheckErr(err)

	rootCmd.PersistentFlags().String("temp", DefaultGorootsDir, "temp directory")
	err = viper.BindPFlag("temp_dir", rootCmd.PersistentFlags().Lookup("temp"))
	cobra.CheckErr(err)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in config directory with name ".godl" (without extension).
		viper.AddConfigPath(defaultConfigPath())
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		cobra.CheckErr(err)
	}

	// SetDefault
	viper.SetDefault("golang_url", DefaultGolangUrl)
	viper.SetDefault("goroots_dir", filepath.Join(defaultConfigPath(), DefaultGorootsDir))
	viper.SetDefault("temp_dir", filepath.Join(defaultConfigPath(), DefaultTempDir))
	viper.SetDefault("versions", []string{})

	cobra.CheckErr(viper.Unmarshal(&cfg))
}

func defaultConfigPath() string {
	config, err := os.UserConfigDir()
	cobra.CheckErr(err)

	if runtime.GOOS == "darwin" {
		config = os.Getenv("XDG_CONFIG_HOME")
		if config == "" {
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)

			if home == "" {
				_, err := fmt.Fprintln(os.Stderr, "$XDG_CONFIG_HOME or $HOME are not defined")
				cobra.CheckErr(err)
			}
			config = filepath.Join(home, ".config")
		}
	}

	return filepath.Join(config, "godl")
}
