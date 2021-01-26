module github.com/lemon-cloud-project/lemon-cloud-service

go 1.15

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/inflection v1.0.0
	github.com/lemon-cloud-project/lemon-cloud-commons-golang v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-project/lemon-cloud-sdk-ext-service v0.0.0-00010101000000-000000000000 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
)

replace github.com/lemon-cloud-project/lemon-cloud-sdk-ext-service => ../lemon-cloud-sdk-ext-service

replace github.com/lemon-cloud-project/lemon-cloud-commons-golang => ../lemon-cloud-commons-golang
