package realnameVierfy

type IdcardVierfy interface {
	Verify(realName, idcard string) (bool, error)
}

//type IdcardVierfy interface {
//	Verify(realName, realnameVierfy string) (bool, error)
//}
func NewIdcardVierfy(use string, secretKeys ...string) IdcardVierfy {
	switch use {
	//蜜堂有信-百度api市场接口
	case "mitangyouxing_ali":
		userLicenseNo := secretKeys[0]
		appcode := secretKeys[1]
		return New_MitangcardVierfy_UseAliApi(userLicenseNo, appcode)
	//蜜堂有信-阿里云api市场接口
	case "mitangyouxing_baidu":
		userLicenseNo := secretKeys[0]
		appcode := secretKeys[1]
		return New_MitangcardVierfy_UseBaiduApi(userLicenseNo, appcode)
	//云亿通-阿里云api市场接口
	case "yunyitong":
		appcode := secretKeys[0]
		return New_YunyitongcardVierfyForTest(appcode)
	//百度api-账号需要企业认证
	case "baiduapi":
		accessKey, secretKey := secretKeys[0], secretKeys[1]
		return New_BaiduIdcardVerify(accessKey, secretKey)
	}
	return nil
}
