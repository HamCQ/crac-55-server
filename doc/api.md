## 🔗基础链接

当前测试地址：https://dev.hamcq.cn  

如果额外说明，接口返回为 `json` 格式


## 数据统计-总数

Method: `GET`

```
/v1/55/analyse/total/2023
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
    "data":{
        "log_num":98196,         //QSO数量
        "award_num":5261,        //奖状数量
        "single_call_num":16357  //独立呼号
    }
}
```