from pyrfc import Connection

from pprint import pprint
import datetime
try:
    from ConfigParser import ConfigParser
except ModuleNotFoundError as e:
    from configparser import ConfigParser

import sys

##实现CALL RFC 多个FM；及返回参数的处理；
##前期配置：参考https://github.com/SAP/PyRFC

imp = dict(
    RFCINT1=0x7f, # INT1: Integer value (1 byte)
    RFCINT2=0x7ffe, # INT2: Integer value (2 bytes)
    RFCINT4=0x7ffffffe, # INT: integer value (4 bytes)
    RFCFLOAT=1.23456789, # FLOAT

    RFCCHAR1=u'a', # CHAR[1]
    RFCCHAR2=u'ij', # CHAR[2]
    RFCCHAR4=u'bcde', # CHAR[4]
    RFCDATA1=u'k'*50, RFCDATA2=u'l'*50, # CHAR[50] each

    RFCTIME=datetime.time(12,34,56), # TIME
    RFCDATE=datetime.date(2012,10,3), # DATE

    RFCHEX3=b'\x66\x67\x68' # BYTE[3]: String with 3 hexadecimal values (='fgh')
)

def connrfc():
    global Conn
    config = ConfigParser()
    config.read('sapnwrfc.cfg')
    params_connection = config._sections['connection']
    try:
        Conn = Connection(**params_connection)
        #list_slice(s) #切片截取后，调用SAP RFC    r =  callRFC(s[start:end])        
        #callRFC()        
        #Conn.close()
    except pyrfc.RFCError as rfcerr:
        print("No error occured",rfcerr)
        Conn = Connection(**params_connection)

def callRFC():
    #连续CALL RFC 多个FM
    result = Conn.call('ZRFC_PPBOM_PLC', WERKS='YK01',GLTRP='20140903',MANTR_MARK='X')#以字段方式传参
    for key in result['IMAKT']:#遍历FM返回字典的每行中某个字段内容
        pprint(key['MAKTX'])
        
    result1 = Conn.call('STFC_STRUCTURE', IMPORTSTRUCT=imp)#以字典方式传参
    pprint(result1)
    
def main():

    connrfc()
    callRFC()
    Conn.close()
    

    # config = ConfigParser()
    # config.read('sapnwrfc.cfg')
    # params_connection = config._sections['connection']

    # conn = Connection(**params_connection)
    #result = conn.call('STFC_STRUCTURE', IMPORTSTRUCT=imp)#以字典方式传参
    #result = conn.call('ZRFC_PPBOM_PLC', WERKS='YK01',GLTRP='20140903',MANTR_MARK='X')#以字段方式传参
    #pprint(result)#打印出所有FM的EXPORT, CHANGING, and TABLE parameters返回内容，以字典方式返回
    #pprint(result['IMAKT'])#指定FM的EXPORT, CHANGING, and TABLE parameters返回某个字典内容
    #for key in result['IMAKT']:#遍历FM返回字典的每行中某个字段内容
    #    pprint(key['MAKTX'])

if __name__ == '__main__':
    main()
