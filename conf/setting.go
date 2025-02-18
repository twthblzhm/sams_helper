package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sams_helper/notice"
)

type ProxySet struct {
	IsEnabled bool   `yaml:"isEnabled"`
	ProxyUrl  string `yaml:"proxyUrl"`
}

type AutoFixPurchaseLimitSet struct {
	IsEnabled  bool `yaml:"isEnabled"`
	FixOffline bool `yaml:"fixOffline"`
	FixOnline  bool `yaml:"fixOnline"`
}

type SupplySet struct {
	Mode      int      `yaml:"mode"`
	IsEnabled bool     `yaml:"isEnabled"`
	KeyWords  []string `yaml:"keyWords"`
}

type SleepTimeSet struct {
	StepStoreSleep    int `yaml:"stepStoreSleep"`
	StepCartSleep     int `yaml:"stepCartSleep"`
	StepCartShowSleep int `yaml:"stepCartShowSleep"`
	StepGoodsSleep    int `yaml:"stepGoodsSleep"`
	StepCapacitySleep int `yaml:"stepCapacitySleep"`
	StepOrderSleep    int `yaml:"stepOrderSleep"`
	StepSupplySleep   int `yaml:"stepSupplySleep"`
}

type AutoInputSet struct {
	IsEnabled       bool  `yaml:"isEnabled"`
	InputPayMethod  int   `yaml:"inputPayMethod"`
	InputAddress    int   `yaml:"inputAddress"`
	InputCouponList []int `yaml:"inputCouponList"`
}

type Setting struct {
	AuthToken               string                  `yaml:"authToken"`
	RunMode                 int                     `yaml:"runMode"`
	SupplySet               SupplySet               `yaml:"supplySet"`
	BruteCapacity           bool                    `yaml:"bruteCapacity"`
	SleepTimeSet            SleepTimeSet            `yaml:"sleepTimeSet"`
	DeviceType              int64                   `yaml:"deviceType"`
	DeliveryType            int64                   `yaml:"deliveryType"`
	StoreType               int64                   `yaml:"storeType"`
	FloorId                 int64                   `yaml:"floorId"`
	IgnoreInvalid           bool                    `yaml:"ignoreInvalid"`
	AutoFixPurchaseLimitSet AutoFixPurchaseLimitSet `yaml:"autoFixPurchaseLimit"`
	PerDateLen              int                     `yaml:"perDateLen"`
	SassId                  string                  `yaml:"sassId"`
	ProxySet                ProxySet                `yaml:"proxySet"`
	NoticeSet               notice.NoticerSet       `yaml:"noticeSet"`
	RunUnlimited            bool                    `yaml:"runUnlimited"`
	AutoInputSet            AutoInputSet            `yaml:"autoInputSet"`
}

func InitSetting() (error, Setting) {
	setting := Setting{}
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err, setting
	}

	err = yaml.Unmarshal(file, &setting)
	if err != nil {
		return err, setting
	}
	return nil, setting
}
