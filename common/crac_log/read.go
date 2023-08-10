package craclog

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadADIFFileLine 从文件中读取adif日志
func ReadADIFFileLine(filename string) ([]string, error) {
	var temp []string
	fi, err := os.Open(filename)
	if err != nil {
		return temp, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	findEOR := ""
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return temp, err
		}
		if strings.Contains(strings.ToLower(string(line)), "eor") {
			findEOR += strings.Replace(string(line), "\n", "", -1)
			temp = append(temp, findEOR)
			findEOR = ""
		} else {
			findEOR += strings.Replace(string(line), "\n", "", -1)
		}

	}
	return temp, nil
}

// ReadAdifString 从adif字符串中解析
func ReadAdifString(adif string) ([]string, error) {
	var temp []string
	br := bufio.NewReader(strings.NewReader(adif))
	findEOR := ""
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return temp, err
		}
		if strings.Contains(strings.ToLower(string(line)), "eor") {
			findEOR += strings.Replace(string(line), "\n", "", -1)
			temp = append(temp, findEOR)
			findEOR = ""
		} else {
			findEOR += strings.Replace(string(line), "\n", "", -1)
		}

	}
	return temp, nil
}
