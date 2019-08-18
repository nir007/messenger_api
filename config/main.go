package config

var Main map[string]map[string]string

func init() {
	Main = map[string]map[string]string{}

	Main["jwt"] = map[string]string{
		"Realm": "sitius",
		"Key": "2koghrhjsdfdfsfwe",
		"Timeout": "100",
		"MaxRefresh": "100",
		"IdentityKey": "id",
	}
}