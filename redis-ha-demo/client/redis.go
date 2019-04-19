package client

import (
	"errors"
	"os/exec"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

const (
	cliPath = "redis-cli"
	netWork = "tcp"
	addr    = "127.0.0.1"
	port    = "6379"
	expire  = 3600 // 有效期(秒)
)

func SetKV(key string, value string) error {
	conn, err := redis.Dial(netWork, addr+":"+port)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, expire)
	return err
}

func GetKV(key string) (string, error) {
	conn, err := redis.Dial(netWork, addr+":"+port)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	res, err := redis.String(conn.Do("GET", key))
	if len(res) == 0 {
		return "", errors.New("not exists")
	}
	return res, nil
}

func Sleep(duraction int) {
	cmd := exec.Command(cliPath, "-h", addr, "-p", port, "DEBUG", "sleep", strconv.Itoa(duraction))
	cmd.Run()
}
