# 基础链接

当前测试地址：https://dev.hamcq.cn  

如果没有额外说明，接口返回均为 `json` 格式

接口预览：  
* 首页搜索 `/v1/55/search`
* 排名统计 - Top5 `/v1/55/analyse/rank/top5`
* 排名统计 - 所有 `/v1/55/analyse/rank/all`
* 数据统计 - 总数 `/v1/55/analyse/total`
* 数据统计 - BY 电台通联数量统计（0-9 区） `/v1/55/analyse/barchart/by09`
* 数据统计 - BnCRA 电台通联统计 / BnCRA 电台通联模式比例 `/v1/55/analyse/barchart/bncra`
* 数据统计 - 省份 QSO 数量统计 `/v1/55/analyse/province`
* Slot 查询 `/v1/55/slot`
* 奖状查询 `/v1/55/award`
* 配置查询 `/v1/55/config`


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
        ],
        "rank_info":{
            "cra": 1,
            "diff_cra": 1,
            "is_cn": false
        }
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

# 奖状查询
Method: `GET`

```
/v1/55/award
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| year | 2022 | String | N | 所属年份 |
| callsign | BH3WNL | String | Y | 呼号 |

Response:  
```
{
    "code":200,
    "status":true,
    "data":{
        "award_string": "Gold",
        "award_type": 1,
        "bncra_num": 10,
        "combination": 27, //组合数
        "callsign": "BH3WNL",
        "continent": "AS", //所在大洲
        "img_url": "", //缩略图
        "img_url_origin": "" //原图
    }
}
```

# 配置查询
Method: `GET`

```
/v1/55/config
```

Request:
| Name | E.g. | Type | Require | Remark |
| --- | --- | --- | --- | --- |
| year | 2023 | String | Y | 所属年份 |

Response:  
```
{
    "code":200,
    "status":true,
    "data":{
        "cover": "", //首页配图
        "crac_desc": "", //crac活动介绍链接
        "current_active_year": 2023, //当前开启活动
        "sub_title": "", //首页标题第二行
        "sub_title_en": "", //英文
        "title": "", //首页标题
        "title_en": ""
    }
}
```