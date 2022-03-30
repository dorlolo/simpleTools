import os
import sqlite3


class AreaReader():
    def __init__(self, path):
        self.conn = sqlite3.Connection(path)
        self.cu = self.conn.cursor()

    # sqlite中不可用
    def initTableInMysql(self):
        self.conn.execute("""DROP TABLE IF EXISTS `sys_dict_area`;""")
        sql = '''
CREATE TABLE `sys_dict_area` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` int(11) DEFAULT NULL COMMENT '父id',
  `short_name` varchar(100) DEFAULT NULL COMMENT '简称',
  `name` varchar(100) DEFAULT NULL COMMENT '名称',
  `merger_name` varchar(255) DEFAULT NULL COMMENT '全称',
  `level` tinyint(4) DEFAULT NULL COMMENT '层级 0 1 2 省市区县',
  `pinyin` varchar(100) DEFAULT NULL COMMENT '拼音',
  `phone_code` varchar(100) DEFAULT NULL COMMENT '长途区号',
  `zip_code` varchar(100) DEFAULT NULL COMMENT '邮编',
  `first` varchar(50) DEFAULT NULL COMMENT '首字母',
  `lng` varchar(100) DEFAULT NULL COMMENT '经度',
  `lat` varchar(100) DEFAULT NULL COMMENT '纬度',
  `area_code` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `ix_china_area_merge` (`merger_name`),
  KEY `ix_china_area_area_code` (`area_code`)
) ;

SET FOREIGN_KEY_CHECKS = 1;
'''
        self.conn.execute(sql)

    def initDbInsqlite(self):
        sqldata = """
                create table sys_dict_area
        (
            pid         integer not null,
            short_name  text    not null,
            merger_name text,
            level       integer,
            pinyin      text,
            phone_code  text,
            zip_code    text,
            first       text,
            lng         text,
            lat         text,
            area_code   text,
            name        text
        );
        create unique index sys_dict_area_id_uindex
            on sys_dict_area (id);

        create unique index sys_dict_area_pid_uindex
            on sys_dict_area (pid);
                """
        self.conn.execute(sqldata)

if __name__ == '__main__':
    # BASE_DIR = os.path.dirname(os.path.abspath(__file__))
    # db_path = os.path.join(BASE_DIR, "area.db")
    areadb = AreaReader("area.db")
    areadb.initDbInsqlite()
    # areadb.cu.execute("""INSERT INTO test(name) VALUES ( ? );""", ("水电费是",))
    # print(areadb.cu.lastrowid)
    # a = areadb.conn.commit()
    # print(a)
