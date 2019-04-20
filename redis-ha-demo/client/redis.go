package client

import (
	"errors"
	"os/exec"
	"strconv"

	"github.com/gomodule/redigo/redis"

	"hackerthon2019/redis-ha/redis-ha-demo/setting"
)

func SetKV(key string, value string) error {
	conn, err := redis.Dial(setting.Config.Redis.Network, setting.Config.Redis.Addr+":"+setting.Config.Redis.Port)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, setting.Config.Redis.Expire)
	return err
}

func GetKV(key string) (string, error) {
	conn, err := redis.Dial(setting.Config.Redis.Network, setting.Config.Redis.Addr+":"+setting.Config.Redis.Port)
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

func Sleep(duration int) {
	cmd := exec.Command(setting.Config.Redis.CliPath, "-h", setting.Config.Redis.Addr, "-p", setting.Config.Redis.Port, "DEBUG", "sleep", strconv.Itoa(duration))
	cmd.Run()
}
