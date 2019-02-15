package config

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strings"
)

func loadCfg() (*map[string]string, error) {
	result := make(map[string]string)

	f, err := os.Open(".env")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		isEof := false
		str, err := br.ReadString('\n')
		str = strings.TrimSpace(str)
		if err != nil {
			if err == io.EOF {
				isEof = true
			} else {
				return nil, err
			}
		}

		if str != "" {
			splitStr := strings.Split(str, "=")
			if len(splitStr) == 2 {
				result[strings.ToUpper(splitStr[0])] = splitStr[1]
			}
		}

		if isEof {
			break
		}
	}

	return &result,nil
}

func saveCfg(env *map[string]string) error {
	f, err := os.OpenFile(".env", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	keyList := make([]string, len(*env))
	index := 0
	for k,_ := range *env {
		keyList[index] = k
		index++
	}
	sort.Strings(keyList)

	for _, k := range keyList {
		_, err := f.WriteString(strings.ToUpper(k) + "=" + (*env)[k] + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}