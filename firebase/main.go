package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/zabawaba99/firego"
)

func main() {
	fmt.Printf("Hello world\n")

	viper.SetConfigName("config") // 設定ファイル名を拡張子抜きで指定する
	viper.AddConfigPath(".")      // 現在のワーキングディレクトリを探索することもできる
	err := viper.ReadInConfig()   // 設定ファイルを探索して読み取る
	if err != nil {               // 設定ファイルの読み取りエラー対応
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}

	f := firego.New(viper.GetString("url"), nil)
	f.Auth(viper.GetString("token"))
	var v map[string]interface{}
	if err := f.Value(&v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", v)
}
