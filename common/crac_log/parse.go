package craclog

import (
	"crac55/common/tools"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Parse 解析日志
func Parse(line []string) ([]CQLog, []ErrInfo, error) {
	var cqlog []CQLog
	var errInfo []ErrInfo
	for k, v := range line {
		if v == "" {
			errInfo = append(errInfo, ErrInfo{k + 1, "存在空行", false})
			continue
		}
		leftCount := strings.Count(v, "<")
		rightCount := strings.Count(v, ">")
		if (leftCount+rightCount)%2 != 0 {
			errInfo = append(errInfo, ErrInfo{k + 1, "格式有误:”<“ 或 ”>” 不匹配", false})
			continue
		}

		compileRegex := regexp.MustCompile("<(.*?)>")
		matchArr := compileRegex.FindAllStringSubmatch(v, -1)
		// fmt.Println(matchArr)
		//[[<CALL:6> CALL:6] [<QSO_DATE:8> QSO_DATE:8] [<TIME_ON:6> TIME_ON:6] [<TIME_OFF:6> TIME_OFF:6] [<BAND:3> BAND:3] [<STATION_CALLSIGN:5> STATION_CALLSIGN:5] [<FREQ:8> FREQ:8] [<CONTEST_ID:2> CONTEST_ID:2] [<FREQ_RX:8> FREQ_RX:8] [<MODE:3> MODE:3] [<RST_RCVD:2> RST_RCVD:2] [<RST_SENT:2> RST_SENT:2] [<TX_PWR:3> TX_PWR:3] [<OPERATOR:5> OPERATOR:5] [<CQZ:2> CQZ:2] [<STX:1> STX:1] [<APP_N1MM_POINTS:1> APP_N1MM_POINTS:1] [<APP_N1MM_RADIO_NR:1> APP_N1MM_RADIO_NR:1] [<APP_N1MM_CONTINENT:2> APP_N1MM_CONTINENT:2] [<APP_N1MM_RUN1RUN2:1> APP_N1MM_RUN1RUN2:1] [<APP_N1MM_RADIOINTERFACED:1> APP_N1MM_RADIOINTERFACED:1] [<APP_N1MM_ISORIGINAL:4> APP_N1MM_ISORIGINAL:4] [<APP_N1MM_NETBIOSNAME:15> APP_N1MM_NETBIOSNAME:15] [<APP_N1MM_ISRUNQSO:1> APP_N1MM_ISRUNQSO:1] [<APP_N1MM_ID:32> APP_N1MM_ID:32] [<APP_N1MM_CLAIMEDQSO:1> APP_N1MM_CLAIMEDQSO:1] [<EOR> EOR]]
		var adfi CQLog
		for _, single := range matchArr {
			//[<RST_RCVD:2> RST_RCVD:2]
			singleLen := len(single)
			if singleLen <= 1 {
				errInfo = append(errInfo, ErrInfo{k + 1, "格式有误", false})
				continue
			}
			s, msg := dealSingle(v, single, &adfi)
			if !s {
				errInfo = append(errInfo, ErrInfo{k + 1, msg, false})
				continue
			}
		}
		cqlog = append(cqlog, adfi)
	}
	return cqlog, errInfo, nil
}

// dealSingle 逐行处理
func dealSingle(line string, match []string, adfi *CQLog) (bool, string) {
	var err error
	if strings.ToLower(match[1]) == "eor" || strings.ToLower(match[1]) == "eoh" {
		return true, ""
	}
	//CALL:6
	temp := strings.Split(match[1], ":")
	if len(temp) != 2 && len(temp) != 3 {
		//兼容带D的情况 <QSO_DATE:8:D>20210504
		return false, "格式有误"
	}
	lower := strings.ToLower(match[1])
	if strings.Contains(lower, "call:") {
		info, err := getTagData(line, match)
		adfi.Call = strings.ToUpper(info)
		if err != nil {
			return false, err.Error()
		}
		bycode, err := tools.GetByCode(info)
		if err != nil {
			return false, "call 格式有误"
		}
		adfi.CNByCode = int8(bycode)
		adfi.CnRegionCode = tools.CateRetionCode(info)
		return true, ""
	}
	if strings.Contains(lower, "mode:") {
		if adfi.Mode != "" && strings.Contains(lower, "submode") {
			submode, err := getTagData(line, match)
			if err != nil {
				return false, err.Error()
			}
			if strings.ToUpper(submode) == "FT4" {
				adfi.Mode = "FT4"
			}
			return true, ""
		}
		//暂时防止3.0的覆盖
		if adfi.Mode != "" {
			return true, ""
		}
		info, err := getTagData(line, match)
		adfi.Mode = strings.ToUpper(info)
		if err != nil {
			return false, err.Error()
		}
		return true, ""
	}
	if strings.Contains(lower, "band:") {
		info, err := getTagData(line, match)
		adfi.Band = strings.ToUpper(info)
		if err != nil {
			return false, err.Error()
		}
		return true, ""
	}
	if strings.Contains(lower, "qso_date:") {
		if len(temp) == 2 {
			adfi.QSODate, err = getTagData(line, match)
			if err != nil {
				return false, err.Error()
			}
		}
		if len(temp) == 3 {
			adfi.QSODate, err = getTagDataWithD(line, match)
			if err != nil {
				return false, err.Error()
			}
		}
		t, err := time.Parse("20060102", adfi.QSODate)
		if err != nil {
			return false, "时间格式有误"
		}
		month := t.Month()
		day := t.Day()
		if month != 5 {
			return false, "QSO 日期需在 5月1日-5日 期间"
		}
		if day > 5 {
			return false, "QSO 日期需在 5月1日-5日 期间"
		}
		return true, ""
	}
	if strings.Contains(lower, "freq") {
		adfi.Frequency, err = getTagData(line, match)
		if err != nil {
			return false, err.Error()
		}
		return true, ""
	}
	if strings.Contains(lower, "station_call") {
		info, err := getTagData(line, match)
		adfi.StationCallsign = strings.ToUpper(info)
		if err != nil {
			return false, err.Error()
		}
		if !tools.IsBnCra(info) {
			return false, "station_callsign 信息必须是 BnCRA"
		}
		return true, ""
	}
	if strings.Contains(lower, "operator") {
		//<operator>后面不能是BnCRA
		info, err := getTagData(line, match)
		adfi.Operator = strings.ToUpper(info)
		if err != nil {
			return false, err.Error()
		}
		if tools.IsBnCra(info) {
			return false, "operator 不能是 BnCRA"
		}
		return true, ""
	}
	if strings.Contains(lower, "time_on") {
		adfi.TimeOn, err = getTagData(line, match)
		if err != nil {
			return false, err.Error()
		}
		return true, ""
	}
	if strings.Contains(lower, "time_off") {
		adfi.TimeOff, err = getTagData(line, match)
		if err != nil {
			return false, err.Error()
		}
		return true, ""
	}
	if strings.Contains(lower, "qsl_rcvd") {
		adfi.QslRcvd, err = getTagData(line, match)
		if err != nil {
			return false, err.Error()
		}
		return true, ""
	}
	return true, ""
}

// getTagData 获取括号数据
func getTagData(line string, matchArray []string) (string, error) {
	//<CALL:6>
	typeIndex := strings.Index(line, matchArray[0])
	//len("<CALL:6>")
	start := typeIndex + len(matchArray[0])
	//:6>
	lenString := strings.Split(matchArray[1], ":")
	//6
	num, err := strconv.Atoi(lenString[1])
	if err != nil {
		return "", errors.New("格式有误")
	}
	end := start + num

	//<STATION_CALLSIGN:6>BG5UWQ<
	compileRegex := regexp.MustCompile(matchArray[0] + "(.*?)<")
	matchArr := compileRegex.FindAllStringSubmatch(line, -1)
	if len(strings.TrimSpace(matchArr[0][1])) != num {
		return line[start:end], errors.New("长度有误:" + matchArray[0] + matchArr[0][1])
	}

	return line[start:end], nil
}

// getTagDataWithD 处理:D的数据
func getTagDataWithD(line string, matchArray []string) (string, error) {
	typeIndex := strings.Index(line, matchArray[0])
	//<QSO_DATE:8:D>20210504
	//len("<QSO_DATE:8:D>")
	start := typeIndex + len(matchArray[0])
	lenString := strings.Split(matchArray[1], ":")
	len, err := strconv.Atoi(lenString[1])
	if err != nil {
		return "", errors.New("格式有误")
	}
	end := start + len
	return line[start:end], nil
}
