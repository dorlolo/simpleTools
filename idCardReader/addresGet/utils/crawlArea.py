import time

import requests
from bs4 import BeautifulSoup

from utils.sqlite import AreaReader
import pathlib
from pypinyin import lazy_pinyin


class Crawler():
    def __init__(self):
        self.wbesite = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021"
        thisPath = pathlib.Path(__file__).parent / "area.db"
        self.db = AreaReader(thisPath.__str__())

    def main(self):
        trs = self.get_response(self.wbesite, 'provincetr')
        try:
            for tr in trs:  # 循环每一行
                datas = []
                for td in tr:  # 循环每个省
                    level = 2  # 1国家 2省 3市 4区
                    province_name = td.a.get_text()
                    province_url = self.wbesite + "/"+td.a.get('href')
                    province_code = tr.find_all('td')[0].string if tr.find_all('td')[0].string else ""#这个是空值
                    print(province_name)

                    # 数据库入库
                    merger_name = f"中国,{province_name}"
                    pinyin = "".join(lazy_pinyin(province_name)).upper()
                    shortName = province_name.replace("自治区","").replace("省","").replace("市","").replace("区","").replace("县","")
                    firstEN = pinyin[:1]
                    sql_insert = '''
                        INSERT INTO
                          sys_dict_area(pid,name,short_name, merger_name, level,pinyin,area_code,first)
                        VALUES
                          (?, ?, ?, ?, ?, ?, ?, ?);
                        '''
                    self.db.cu.execute(sql_insert, (
                    1, province_name, shortName, merger_name, level, pinyin, province_code, firstEN,))
                    province_id = self.db.cu.lastrowid
                    self.db.conn.commit()

                    trs = self.get_response(province_url, None)
                    for tr in trs[1:]:  # 循环每个市
                        city_code = tr.find_all('td')[0].string
                        city_name = tr.find_all('td')[1].string
                        city_url = self.wbesite +"/" +tr.find_all('td')[1].a.get('href')
                        trs = self.get_response(city_url, None)

                        # 数据入库
                        level = 3
                        merger_name = f"中国,{province_name}-{city_name}"
                        pinyin = "".join(lazy_pinyin(city_name)).upper()
                        shortName = city_name if city_name=="市辖区" else city_name.replace("自治区","").replace("省","").replace("市","").replace("区","").replace("县","")
                        firstEN = pinyin[:1]
                        sql_insert = '''
                            INSERT INTO
                              sys_dict_area(pid,name,short_name, merger_name, level,pinyin,area_code,first)
                            VALUES
                              (?, ?, ?, ?, ?, ?, ?, ?);
                            '''
                        self.db.cu.execute(sql_insert, (
                        province_id, city_name, shortName, merger_name, level, pinyin, city_code, firstEN,))
                        city_id = self.db.cu.lastrowid
                        self.db.conn.commit()

                        for tr in trs[1:]:  # 循环每个区
                            time.sleep(0.5)
                            county_code = tr.find_all('td')[0].string
                            county_name = tr.find_all('td')[1].string
                            # data = [province_name, city_code, city_name, county_code, county_name]
                            # print(data)
                            # datas.append(data)
                            # 数据入库
                            level = 4
                            merger_name = f"中国,{province_name}-{city_name}-{county_name}"
                            pinyin = "".join(lazy_pinyin(county_name)).upper()
                            shortName = county_name.replace("自治区","").replace("省","").replace("市","").replace("区","").replace("县","")
                            firstEN = pinyin[:1]
                            sql_insert = '''
                                INSERT INTO
                                  sys_dict_area(pid,name,short_name, merger_name, level,pinyin,area_code,first)
                                VALUES
                                  (?, ?, ?, ?, ?,  ?, ?, ?);
                                '''
                            self.db.cu.execute(sql_insert, (city_id, county_name, shortName, merger_name, level, pinyin, county_code, firstEN))
                            self.db.conn.commit()

                        time.sleep(2)
                # sql = "insert into china (province_name,city_code,city_name,county_code,county_name) values (%s,%s,%s,%s,%s)"
                # self.connect_mysql(sql, datas)
        except Exception as e:
            print(e)
            pass

    def get_response(self, url, attr):
        """
        ————————————————
        版权声明：本文为CSDN博主「CoderYYN」的原创文章，遵循CC
        4.0
        BY - SA版权协议，转载请附上原文出处链接及本声明。
        原文链接：https: // blog.csdn.net / ychgyyn / article / details / 90514998
        :param self:
        :param url:
        :param attr:
        :return:
        """
        response = requests.get(url)
        response.encoding = 'utf-8'  # 编码转换
        soup = BeautifulSoup(response.text, 'lxml')
        table = soup.find_all('tbody')[1].tbody.tbody.table
        if attr:
            trs = table.find_all('tr', attrs={'class': attr})
        else:
            trs = table.find_all('tr')
        return trs


if __name__ == '__main__':
    Crawler().main()
