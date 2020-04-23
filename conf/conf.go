/**
 * conf
 * Created by tiger on Sun.Feb.2020
 */
package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//profile variables
type conf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}
func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func main() {
	var c conf
	conf:=c.getConf()
	fmt.Println(conf.Host)
}
