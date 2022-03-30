from pypinyin import lazy_pinyin

if __name__ == '__main__':
    name="".join(lazy_pinyin("北京市")).upper()
    print(name)