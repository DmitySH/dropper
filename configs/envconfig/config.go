package envconfig

import (
	"embed"
	"github.com/spf13/viper"
	"log"
)

//go:embed app.env
var embEnv embed.FS

func LoadEnvConfig(cfgPath string) {
	viper.SetConfigFile(cfgPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("can't load env config:", err)
		log.Println("using system envs")
	}
}

func LoadEmbeddedEnvConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	f, openErr := embEnv.Open("app.env")
	if openErr != nil {
		log.Fatal("can't read embedded config file")
	}
	defer f.Close()

	if err := viper.ReadConfig(f); err != nil {
		log.Println("can't load env config:", err)
		log.Println("using system envs")
	}
}
