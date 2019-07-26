#-*- coding-8 -*-
import requests
import lxml
import sys
from bs4 import BeautifulSoup
import xlwt
import time
import urllib
import json

headers = {
            'Authorization': 'efe36d46-dc2a-4228-8d3b-eeb5daa5eaf1',
           }
try:
    response = requests.get('http://open.api.tianyancha.com/services/open/ic/baseinfo/2.0?name=大成食品（大连）有限公司',headers = headers)
    if response.status_code != 200:
        response.encoding = 'utf-8'
        print(response.status_code)
        print('ERROR')
    soup = BeautifulSoup(response.text,'lxml')
    #print(soup.html.body.p)
    #print(soup.find_all('p'))#打到所有P标签下内容
#这是HTML返回格式：<html><body><p>{"result":{"staffNumRange":null,

    json1 = soup.find_all('p')[0].get_text()  # 获取p标签文本内容
    #print(json1)
	
	# 将 JSON 对象转换为 Python 字典
	#data2 = json.loads(json_str)
	#print ("data2['name']: ", data2['name'])
   
    jsonobj = json.loads(json1) #格式：{"result":{"staffNumRange":null,
    #转字典格式
    data1 = jsonobj['result'] #格式{"staffNumRange":null,"revokeReason":null,"fromTime":818179200000}
    print(data1['name'])
    print(data1['regLocation'],data1['taxNumber'],data1['legalPersonName'])



except Exception:
        print('请求都不让，这天眼查也搞事情吗？？？')