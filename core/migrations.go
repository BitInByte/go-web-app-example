package core

import "github.com/BitInByte/web-app-example/model"


func Migrations()  {
    DB.AutoMigrate(&model.User{})
}
