
from venv import logger
from pyrfc import Connection
import pyrfc
from pprint import pprint
try:
    from ConfigParser import ConfigParser
except ModuleNotFoundError as e:
    from configparser import ConfigParser

import csv
import os
import sys
import datetime
import threading
import time
import shutil
'''
功能：读取某目录下所有CSV UTF8文件，调取SAP RFC并存入SAP DB 表中，最后把CSV文件+时间戳移到其它目录。
部署准备：1.指定目录：file_dir；2.指定备份目录：new_file_path
'''

#日期格式转换
def strday(s):
    d = datetime.datetime.strptime(s, "%d/%m/%Y")
    strd = d.strftime('%Y%m%d')
    # 03/04/2019 to 20190412 日期格式转换
    return strd

def fun_timer():
# ##读取目录下所有文件，以列表形式存取并切片，在切片数据后调用SAP RFC
    config = ConfigParser()
    config.read('sapnwrfc.cfg')
    params_filelist = config._sections['filelist']
    #print(params_filelist['file_dir'])
    
    #file_dir = 'D:\\test'
    file_dir = params_filelist['file_dir']
    new_file_path = params_filelist['backup_dir']
    L = file_name(file_dir)  # 读取目录下所有CSV文件，以列表形式
    print(L)
    for csvf in L:
        print(csvf)
        list_csv = opencsv(csvf) #以列表形式存取并切片
        connrfc(list_csv)#连接SAP后，再切片数据后调用SAP RFC
        #开始剪切CSV文件至指定目录new_file_path 新目录不能放在file_path目录下（避免重复导入）
        ##开始给文件名加时间戳        
        len_csvf = len(csvf)
        lstr = len_csvf - 4
        t = time.time()
        timefilepath = csvf[:lstr]+"_"+str(datetime.datetime.now().strftime('%Y-%m-%d %H%M%S'))+".csv"
        print(timefilepath)
        os.rename(csvf,timefilepath)
        file_path = timefilepath
        ##开始给文件名加时间戳
        #file_path = csvf
        #new_file_path = 'D:\\testbackcsv'
        try:
            shutil.move(file_path, new_file_path)
            logger.info("{a}条剪切成功！".format(a=i))
        except Exception as e:
            logger.info("剪切失败：{a}".format(a=e))

# ##读取目录下所有文件，以列表形式
    print('hello timer')  # 打印输出
    global timer  # 定义变量
    timer = threading.Timer(60, fun_timer)  # 60秒调用一次函数
    # 定时器构造函数主要有2个参数，第一个参数为时间，第二个参数为函数名
    timer.start()  # 启用定时器
    # while True:
    #     time.sleep(1)
    #     print( 'main running......')

#目录下文件获取－只读取CSV文件
def file_name(file_dir):
    L=[]
    for root,dirs,files in os.walk(file_dir):
        #print(root)
        #print(dirs)
        #print(files)
        for file in files:
            if os.path.splitext(file)[1] == '.csv':
                L.append(os.path.join(root,file))
    return L

#打开CSV并存入 列表中
def opencsv(filedir):
    List_csv = []
    # List_csv.clear()
    with open(filedir, newline='',encoding='utf-8') as csvfile:
        spamreader = csv.DictReader(csvfile, delimiter=',', quotechar='"')
        for row in spamreader:
            if row['Dr / Cr Indicator'] == 'Credit' and row['Transaction Details10'] != '' and row['Account Number'] == '501510687378':
                # print(row['Value Date'], row['Transaction Details1'], row['Account Number'],row['Transaction Details10'], i)
                a = row['Transaction Amount']
                jiaoyi_jiner = float(a.replace(',','')) #交易金额格式转换
                t = row['Value Date']
                jaoyi_day = strday(t)	#交易日期格式转换
                #print(jaoyi_day, row['Transaction Details1'], row['Account Number'],row['Transaction Details10'], jiaoyi_jiner)
                #CI17011901159839(流水) 15/01/2019（日期） 501510687378（企业账号）Credit(借贷) 西安鑫德食品原料有限公司（打款名） 23630.0（金额） 货款（附言） 交通银行股份有限公司陕西省分行（开户行名） 611301091018000136879 2019011548995565

                csv_l = [row['Account Number'],row['Transaction Details1'],jaoyi_day,row['Dr / Cr Indicator'],row['Transaction Details10'],jiaoyi_jiner,row['Transaction Details11'],row['Transaction Details12'],row['Transaction Details14'],row['Transaction Details15']]#列表添加
                List_csv.extend(csv_l)
    return List_csv

def connrfc(s):
    global Conn
    config = ConfigParser()
    config.read('sapnwrfc.cfg')
    params_connection = config._sections['connection']
    try:
        Conn = Connection(**params_connection)
        list_slice(s)
        Conn.close()
    except pyrfc.RFCError as rfcerr:
        print("No error occured",rfcerr)
        Conn = Connection(**params_connection)
        
#//分割slice
def list_slice(s):
    
    i = 10
    start = 0
    end = 0
    con = len(s) / i
    c  = 1
    while c <= con:
        print("循环次数：", c)
        sum = 0
        sum += c * i
        start = sum - i
        end = sum
        #print(s[start:end])
        c = c + 1
        #切片截取后，调用SAP RFC
        r =  callRFC(s[start:end])
        print('test_info:',r)
#如果r['RETURN_FLAG'] == 'E'不成功时，是否要考虑记账log本地文件
    
    # return s[start:end]
#CALL SAP RFC
def callRFC(s):
    #CI17011901159839(流水) 15/01/2019（日期） 501510687378（企业账号）Credit(借贷)  西安鑫德食品原料有限公司（打款名） 23630.0（金额） 货款（附言）
    #交通银行股份有限公司陕西省分行（开户行名） 611301091018000136879 2019011548995565

    #RFC传数给值
    imp = dict(
        ACNO=s[0],
        SERIALNO=s[1],
        JIAOYIRQ=s[2],
        SHOUZBS=s[3],
        DFHUMING=s[4],
        JYEDU=s[5],
        FUYAN=s[6],
        DFKHMING=s[7],
        DFZHANGHAO=s[8]
               )
    result = Conn.call('Z_YQZL_SFK01', R_BANKSAP=imp)
    if result['RETURN_FLAG'] == 'E':
        sep = ','
        #sep='\n' #就是分行输入
        t = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        strtime = str(t)
        msg = str(result['RETURN_MSG'])
        listlog = [str(s),msg,strtime,'\n'] #列表 +'\n' 这是换行符号
        fl = open('.\\log\\listlog.txt', 'a')#python需要追加写入文件的时候，是用这个方法f = open('md5_value.txt', 'a')，f = open('md5_value.txt', 'w')这个是不追加写入
        #fl.write(str(listlog))
        fl.write(sep.join(listlog))
        fl.close()
        
        #sep = ','
        #fl.write(sep.join(list))
    
    # with Connection(**params_connection) as conn:
    #     result = conn.call('Z_WUSJ_TEST02',R_CDPIAOJUBS=imp) 
    return result
        #conn.close()
    # conn = Connection(**params_connection)
    # result = conn.call('Z_WUSJ_TEST02',R_CDPIAOJUBS=imp)
    #return result

    ##PS：shutil.move() # 剪切文件；shutil.copy() # 拷贝文件
    ##目录格式如下
    #bank_shoukuan_path = 'D:\\test\\'
    # new_path = 'D:\\'
def zip_file(old_path,new_path):
    # if get_file_count(old_path) >= 10:
        for root, dirs, files in os.walk(old_path):
            for i in range(len(files)):
                # print(files[i])
                if (files[i][-4:] == '.csv'):#or (files[i][-5:] == '.html')
                    file_path = old_path + files[i]
                    # print(file_path)
                    new_file_path = new_path  #+ files[i]
                    # print(new_file_path)
                    try:
                        shutil.move(file_path, new_file_path)
                        logger.info("{a}条剪切成功！".format(a=i))
                    except Exception as e:
                        logger.info("剪切失败：{a}".format(a=e))
    # else:
    #     print("文件少于10条，不需要剪切！")

def main():
    # config = ConfigParser()
    # config.read('sapnwrfc.cfg')
    # params_connection = config._sections['connection']
    # Conn = Connection(**params_connection)
    #主程序定时器
    fun_timer()
'''
    #功能概览：读取某目录下所有CSV UTF8文件，调取SAP RFC并存入SAP DB 表中，最后把CSV文件+时间戳移到其它目录。
    #处理过程：1.读取目录下所有CSV文件，以列表形式存取，连接SAP服务，对切片数据分割后调用SAP RFC，关闭SAP连接。
               最后把CSV文件+时间戳移到其它目录；
    #
'''
   

if __name__ == '__main__':
    main()
