package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"hackerthon2019/redis-ha/redis-ha-demo/client"
)

type kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func set(c *gin.Context) {
	var form kv
	if err := binding.JSON.Bind(c.Request, &form); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	if err := client.SetKV(form.Key, form.Value); err != nil {
		c.JSON(http.StatusBadRequest, "failed to set to redis")
		return
	}

	c.JSON(http.StatusOK, "success")
}

func get(c *gin.Context) {
	var form kv
	if err := binding.JSON.Bind(c.Request, &form); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	res, err := client.GetKV(form.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	c.JSON(http.StatusOK, res)
}

func attack(c *gin.Context) {
	var form struct {
		Duration int `json:"duration"`
	}

	if err := binding.JSON.Bind(c.Request, &form); err != nil || form.Duration < 1 {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	go client.Sleep(form.Duration)
	c.JSON(http.StatusOK, "success")
}

func Init(gs *gin.Engine) {
	gs.GET("/", index)
	gs.POST("/set", set)
	gs.POST("/get", get)
	gs.POST("/attack", attack)
}
