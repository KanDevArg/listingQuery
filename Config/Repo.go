package Config

import "fmt"

type ExtSrcProvider struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}


func BuildExtSrcProvider() *ExtSrcProvider {
	extSrcConfig := ExtSrcProvider{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     "root",
		DBName:   "todos",
		Password: "rootadmin",
	}
	return &extSrcConfig
}


func GetExternalConfigData(extSrcProvider *ExtSrcProvider) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		extSrcProvider.User,
		extSrcProvider.Password,
		extSrcProvider.Host,
		extSrcProvider.Port,
		extSrcProvider.DBName,
	)
}


//Go client
//https://github.com/elastic/go-elasticsearch


// curl -XGET "http://elasticsearch:9200/_search" -H 'Content-Type: application/json' -d'{  "query": {    "match_all": {}  }}'

//docker-compose up -d
//docker-compose down
