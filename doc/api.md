## 🔗基础链接

当前测试地址：https://dev.hamcq.cn  

如果额外说明，接口返回为 `json` 格式

## 首页搜索

Method: `GET`

```
/v1/55/search
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| callsign | BH7CSA | String | Y | 呼号 |
| year | 2023 | String | N | 所属年份 |


Response:  
```
{
    "code":200,
    "status":true,
    "message": "SUCCESS",
    "data":{
        //奖状状态
        "award_info":{
            "award_string": "Gold", //1:Gold 2:Silver 3:Bronze
            "award_type": 1,
            "continent": "AS", //亚洲区域
            "status": true, //是否获奖
        },       
        "bncra": [
            "callsign_station": "BnCRA",
            "cw": {
                "6M": [
                    {
                        "oprator": "",
                        "frequecy": "",
                    }
                ]
            },

        ],       
        "bxcra":[
            {
                "callsign_station": "BnCRA",
                "cw": {
                    "6M": [
                        {
                            "oprator": "",
                            "frequecy": "",
                        }
                    ]
                },
            }
        ] 
    }
}
```

## 排名统计 - Top5


Method: `GET`

```
/analyse/rank/top5
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| year | 2023 | String | N | 所属年份 |


## 排名统计 - 所有


Method: `GET`

```
/analyse/rank/all
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| year | 2023 | String | N | 所属年份 |
| page | 1 | String | N | 页码，每页 100 条 |
| type | 1 | String | N | 类型 |

type:  

```
CnDiffCra     = 1
CnCra         = 2
GlobleDiffCra = 3
GlobleCra     = 4
```

## 数据统计-总数

Method: `GET`

```
/v1/55/analyse/total
```

Request:
```
```

Response:  
```
{
    "code":200,
    "status":true,
    "data":{
        "log_num":98196,         //QSO数量
        "award_num":5261,        //奖状数量
        "single_call_num":16357  //独立呼号
    }
}
```

