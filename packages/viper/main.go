package main

import (
	"fmt"

	gt "github.com/ehsanx64/go-tools/general"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot read config file: %w", err))
	}

	hisAge := viper.GetInt("age")

	gt.SetColumnWidth(10)
	gt.DumpValue("Name", viper.Get("name"))
	gt.DumpValue("Age", hisAge)

	if hisAge == 0 {
		if viper.IsSet("age") {
			fmt.Println("His age is really zero ...")
		} else {
			fmt.Println("His age is not specified at all ...")
		}
	}
}
