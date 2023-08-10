package tools

import (
	"regexp"
	"strings"
)

// CateRetionCode 归类呼号所属省份、自治区、直辖市
func CateRetionCode(callsign string) int {
	//异地台等无法准确得知所在省份
	if strings.Contains(callsign, "/") {
		return 0
	}
	//是否中国大陆/内地呼号
	if !strings.HasPrefix(callsign, "B") {
		return 0
	}
	callLen := len(callsign)
	callsign = strings.ToLower(callsign)

	//特设台等无法准确得知所在省份
	if callLen != 5 && callLen != 6 {
		return 0
	}
	if checkBnCRA(callsign) {
		return 0
	}
	//提取所在区号
	var valid = regexp.MustCompile("[0-9]")
	areaCode := valid.FindString(callsign)
	if len(areaCode) != 1 {
		return 0
	}
	switch areaCode {
	case "0":
		return area0(callsign, callLen)
	case "1":
		return 11
	case "2":
		return area2(callsign, callLen)
	case "3":
		return area3(callsign, callLen)
	case "4":
		return area4(callsign, callLen)
	case "5":
		return area5(callsign, callLen)
	case "6":
		return area6(callsign, callLen)
	case "7":
		return area7(callsign, callLen)
	case "8":
		return area8(callsign, callLen)
	case "9":
		return area9(callsign, callLen)
	}

	return 0
}

// area0 0区校验 新疆 西藏
func area0(callsign string, len int) int {
	code := callsign[3:4]
	//AA~FZ 新疆
	if code >= "a" && code <= "f" {
		return 65
	}
	//GA~LZ 西藏
	if code >= "g" && code <= "l" {
		return 54
	}
	return 0
}

// area2 2区校验 黑龙江 吉林 辽宁
func area2(callsign string, len int) int {
	code := callsign[3:4]
	//AA~HZ 黑龙江
	if code >= "a" && code <= "h" {
		return 23
	}
	//IA~PZ 吉林
	if code >= "i" && code <= "p" {
		return 22
	}
	//QA~XZ 辽宁
	if code >= "q" && code <= "x" {
		return 21
	}
	return 0
}

// area3 3区 天津 内蒙古 河北 山西
func area3(callsign string, len int) int {
	code := callsign[3:4]
	//AA~FZ 天津
	if code >= "a" && code <= "f" {
		return 12
	}
	//GA~LZ 内蒙古
	if code >= "g" && code <= "l" {
		return 15
	}
	//河北
	if code >= "m" && code <= "r" {
		return 13
	}
	//山西
	if code >= "s" && code <= "x" {
		return 14
	}
	return 0
}

// area4 4区 上海 山东 江苏
func area4(callsign string, len int) int {
	code := callsign[3:4]
	//AA~HZ 上海
	if code >= "a" && code <= "h" {
		return 31
	}
	//山东
	if code >= "i" && code <= "p" {
		return 37
	}
	// 江苏
	if code >= "q" && code <= "x" {
		return 32
	}
	return 0
}

// area5 5区 浙江 江西 福建
func area5(callsign string, len int) int {
	code := callsign[3:4]
	//AA~HZ 浙江
	if code >= "a" && code <= "h" {
		return 33
	} //江西
	if code >= "i" && code <= "p" {
		return 36
	}
	// 福建
	if code >= "q" && code <= "x" {
		return 35
	}
	return 0
}

// area6 6区 安徽 河南 湖北
func area6(callsign string, len int) int {
	code := callsign[3:4]
	//AA~HZ 安徽
	if code >= "a" && code <= "h" {
		return 34
	}
	//河南
	if code >= "i" && code <= "p" {
		return 41
	}
	// 湖北
	if code >= "q" && code <= "x" {
		return 42
	}
	return 0
}

// area7 7区 湖南 广东 广西 海南
func area7(callsign string, len int) int {
	code := callsign[3:4]
	//AA~HZ 湖南
	if code >= "a" && code <= "h" {
		return 43
	}
	//广东
	if code >= "i" && code <= "p" {
		return 44
	}
	//广西
	if code >= "q" && code <= "x" {
		return 45
	}
	//海南
	if code >= "y" && code <= "z" {
		return 46
	}
	return 0
}

// area8 四川 重庆 贵州 云南
func area8(callsign string, len int) int {
	code := callsign[3:4]
	//四川
	if code >= "a" && code <= "f" {
		return 51
	}
	//重庆
	if code >= "g" && code <= "l" {
		return 50
	}
	//贵州
	if code >= "m" && code <= "r" {
		return 52
	}
	//云南
	if code >= "s" && code <= "x" {
		return 53
	}
	return 0
}

// area9 陕西 甘肃 宁夏 青海
func area9(callsign string, len int) int {
	code := callsign[3:4]
	//陕西
	if code >= "a" && code <= "f" {
		return 61
	}
	//甘肃
	if code >= "g" && code <= "l" {
		return 62
	}
	//宁夏
	if code >= "m" && code <= "r" {
		return 64
	}
	//青海
	if code >= "s" && code <= "x" {
		return 63
	}
	return 0
}

func checkBnCRA(station string) bool {
	s := []string{
		"B0CRA", "B1CRA", "B2CRA", "B3CRA", "B4CRA",
		"B5CRA", "B6CRA", "B7CRA", "B8CRA", "B9CRA",
	}
	for _, v := range s {
		if station == v {
			return true
		}
	}
	return false
}
