# 基础链接

当前测试地址：https://dev.hamcq.cn  

如果没有额外说明，接口返回均为 `json` 格式

# 首页搜索

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

# 排名统计 - Top5


Method: `GET`

```
/v1/55/analyse/rank/top5
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| year | 2023 | String | N | 所属年份 |


# 排名统计 - 所有


Method: `GET`

```
/v1/55/analyse/rank/all
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

# 数据统计 - 总数

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



# 数据统计 - BY 电台通联数量统计（0-9 区）

Method: `GET`

```
/v1/55/analyse/barchart/by09
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| year | 2023 | String | N | 所属年份 |

Response:  
```
{
    "code":200,
    "status":true,
    "data":[
        {
            "by_code": 5,
            "cw_num": 0,
            "phone_num": 0,
            "digi_num": 0
        }
    ]
}
```

# 数据统计 - BnCRA 电台通联统计 / BnCRA 电台通联模式比例

Method: `GET`

```
/v1/55/analyse/barchart/bncra
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| year | 2023 | String | N | 所属年份 |

Response:  
```
{
    "code":200,
    "status":true,
    "data":[
        {
            "callsign_station": "B0CRA",
            "cw_num": 0,
            "phone_num": 0,
            "digi_num": 0,
            "sum": 0
        }
    ]
}
```


# 数据统计 - 省份 QSO 数量统计

Method: `GET`

```
/v1/55/analyse/province
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| year | 2023 | String | N | 所属年份 |

Response:  
```
{
    "code":200,
    "status":true,
    "data":[
        {
            "cn_region_code": "35",
            "qso_num": 0
        }
    ]
}
```

# Slot 查询

Method: `GET`

```
/v1/55/slot
```

Request:
```
{}
```

Response:  
```
{
    "code":200,
    "status":true,
    "data":[
            "callsign_station": "BnCRA",
            "cw": {
                "6M": [
                    {
                        "callsign": "", //操作员呼号
                        "time": "",     //slot时间
                    }
                ]
            },

    ]
}
```