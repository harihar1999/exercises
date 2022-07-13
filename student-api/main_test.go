package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())


}

func TestGetUsers(t *testing.T){
	response,_ := http.Get("http:localhost:8080/student-api/user")
	fmt.Println(response)
	if response ==nil {
		t.Errorf("got nil want some response")
	}

}




