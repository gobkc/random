package random

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

var ErrorSupport = errors.New("ErrorOnlyAllowSliceOrStruct")
var ErrorSliceMakeFirst = errors.New("ErrorPleaseMakeSliceFirst")
var names = []string{"王", "李", "张", "刘", "陈", "杨", "黄", "吴", "赵", "周", "徐", "孙", "马", "朱", "胡", "林", "郭", "何", "高", "罗", "郑", "梁", "谢", "宋", "唐", "许", "邓", "冯", "韩", "曹", "曾", "彭", "萧", "蔡", "潘", "田", "董", "袁", "于", "余", "叶", "蒋", "杜", "苏", "魏", "程", "吕", "丁", "沈", "任", "姚", "卢", "傅", "钟", "姜", "崔", "谭", "廖", "范", "汪", "陆", "金", "石", "戴", "贾", "韦", "夏", "邱", "方", "侯", "邹", "熊", "孟", "秦", "白", "江", "阎", "薛", "尹", "段", "雷", "黎", "史", "龙", "陶", "贺", "顾", "毛", "郝", "龚", "邵", "万", "钱", "严", "赖", "覃", "洪", "武", "莫", "孔"}
var namesLen = len(names)
var fullName = []string{"威泓", "润棋", "豪霆", "泽军", "鼎曜", "勇瑜", "尚锐", "哲悦", "成嘉", "晟棋", "鹏誉", "鸿俊", "博超", "岸语", "轩拓", "敬荣", "棋浚", "轩俊", "浩荣", "浚才", "翔东", "昊维", "智霖", "浩杰", "宇杰", "思哲", "璟新", "宇文", "智浩", "皓忠", "振东", "皓峰", "璟烨", "子恒", "梓豪", "佩璟", "清风", "皓玮", "韵安", "博平", "志平", "悦晨", "佩玮", "韵林", "佩毅", "豪元", "泰安", "梓轩", "锐博", "方迪", "宏裕", "新铭", "锦骐", "宇梵", "文宣", "宇渤", "琦涛", "志绮", "亚楠", "玉峰", "项哲", "俊良", "靖平", "衡岳", "博易", "琪维", "梦文", "成和", "孟华", "善福", "吉钰", "德惠", "博瀚", "清杰", "德泽", "成文", "博学", "承志", "胜俊", "高畅", "德华", "德庸", "景盛", "成益", "成仁", "波鸿", "德辉", "彭渝", "德明", "铭忠", "陆涛", "华卫", "俊基", "承德", "德宇", "鹏吉", "承泽", "尚凯", "才英", "博艺", "驰勋", "承天", "博文", "承业", "淇峰", "庆元", "德馨", "梦亿", "儒湖", "才俊", "儒林", "凯嘉", "波光", "杰科", "经纶", "德佑", "济琛", "圣尧", "承弼", "博涛", "博延", "煜玮", "思翰", "浩霖", "皓忠", "振东", "皓璟", "谊华", "煦忠", "志平", "洋珈", "悦晨", "皓睿", "泽辉", "文林", "志元", "汕睿", "凯昕", "宇铭", "皓予", "元铭", "熙鹏", "绍峰", "子轩", "智新", "佩玮", "韵林", "湛豪", "嘉逸", "怀冬", "万辰", "知昆", "瑞木", "智崇", "瑾玉", "翰城", "彰苛", "影致", "锦顾", "柏霖", "浩洋", "清霜", "文霖", "玉钦", "志峻", "安境", "尘傲", "何深", "邦溪", "何清", "瀚宁", "铭阔", "泽新", "新释", "叶祺", "含昭", "旭宇", "木琪", "云景", "渊致", "文强", "昕杰", "鸿光", "皓涛", "霖旭", "伦博", "骏杰", "雨泽", "景霖", "嘉诺", "浩堇", "彬畅", "致轩", "晟德", "轩霖", "志铭", "弘鸿", "金然", "星文", "鸿煊", "智鹏", "明硕", "骏志", "展羽", "启天", "雅蕊", "月荷", "珺瑶", "慧馨", "雪青", "茵睿", "初夏", "夜柳", "涵瑶", "竹琼", "岚盈", "月素", "珊花", "涵艳", "欣雪", "琪裳", "昕云", "春瑶", "佩云", "巧艺", "艳秋", "尔曼", "珊琦", "雪春", "歆华", "若美", "韵蓝", "珊叶", "醉薇", "绮彤", "玉冰", "霜瑜", "水宁", "青瑜", "秋翠", "洁云", "欣琪", "欣芳", "海亦", "清惠", "紫青", "雅昕", "佩雅", "雪妍", "白韵", "晓兰", "瑶春", "怜菡", "雅蕾", "琦美", "妙华", "嘉丹", "燕歆", "柳娇", "钰裳", "雪曦", "凤怡", "沐妮", "晓恬", "依芊", "欣芙", "嘉云", "婉嫣", "宛珊", "沐忆", "映佳", "沐芙", "欣黛", "语惠", "依娜", "宛蕊", "姗翾", "宛璇", "蕊媚", "欣芷", "佳芷", "宛玥", "若华", "宛翾", "妮霜", "依欣", "柔宛", "沐娟", "柔依", "雨真", "欣霜", "姗黛", "玥彤", "语雅", "佳忆", "宛盈", "若语", "欣佳", "沐宸", "昕芙", "沐璐", "青桂", "沐玲", "青娜", "宛瑶", "妮慧"}
var fullNameLen = len(fullName)
var sex = []string{"男", "女"}
var phone = []string{"135", "136", "137", "138", "139", "150", "151", "152", "157", "158", "159", "187", "188", "147", "130", "131", "132", "156", "186", "145", "133", "153", "189"}
var phoneLen = len(phone)

var unmarshalMap = map[reflect.Kind]func(dest any) error{
	reflect.Slice: func(dest any) error {
		destType := reflect.Indirect(reflect.ValueOf(dest)).Type()
		meta := reflect.New(destType.Elem()).Interface()

		valueOf := reflect.ValueOf(dest)
		if valueOf.Type().Kind() == reflect.Pointer {
			valueOf = valueOf.Elem()
		}
		length := valueOf.Len()
		if length == 0 {
			return ErrorSliceMakeFirst
		}
		for curRow := 0; curRow < length; curRow++ {
			v := meta
			if err := unmarshalStruct(v); err != nil {
				return err
			}
			valueOf.Index(curRow).Set(reflect.ValueOf(v).Elem())
		}
		return nil
	},
	reflect.Struct: func(dest any) error {
		return unmarshalStruct(dest)
	},
}

var bindRules = map[string]string{
	"username": `(?i)(user|name)`,
	"sex":      `(?i)(sex)`,
	"email":    `(?i)(mail|email)`,
	"phone":    `(?i)(mobile|phone|tel|contact)`,
	"string":   ``, //other string、int、float、float
	"float":    ``,
	"bool":     ``,
}

func Unmarshal(bindSliceOrStruct any) error {
	typeOf := reflect.TypeOf(bindSliceOrStruct)
	if typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
	}
	if process, ok := unmarshalMap[typeOf.Kind()]; ok {
		return process(bindSliceOrStruct)
	}
	return ErrorSupport
}

func genField[T int | int32 | int64 | string | float64 | float32 | bool](field string) (t T) {
	switch field {
	case "username":
		n1 := Number[int](namesLen)
		n2 := Number[int](fullNameLen)
		username := names[n1] + fullName[n2]
		reflect.ValueOf(&t).Elem().SetString(username)
	case "sex":
		s := Number[int](1)
		sexStr := sex[s]
		reflect.ValueOf(&t).Elem().SetString(sexStr)
	case "email":
		email := fmt.Sprintf(`%s@%s.%s`, AlphaNumber(10), AlphaNumber(10), AlphaNumber(10))
		reflect.ValueOf(&t).Elem().SetString(email)
	case "phone":
		n1 := Number[int](phoneLen - 1)
		phoneStr := fmt.Sprintf(`%s%08d`, phone[n1], Number[int](10000000))
		reflect.ValueOf(&t).Elem().SetString(phoneStr)
	case "string":
		reflect.ValueOf(&t).Elem().SetString(AlphaNumber(10))
	case "number":
		n := Number[int64](1000000000)
		numKind := reflect.TypeOf(t).Kind()
		if numKind == reflect.Int {
			reflect.ValueOf(&t).Elem().Set(reflect.ValueOf(int(n)))
		}
		if numKind == reflect.Int32 {
			reflect.ValueOf(&t).Elem().Set(reflect.ValueOf(int32(n)))
		}
		if numKind == reflect.Int64 {
			reflect.ValueOf(&t).Elem().Set(reflect.ValueOf(int64(n)))
		}
	case "float":
		n := Number[int64](15)
		if n < 1000 {
			n = 1000
		}
		reflect.ValueOf(&t).Elem().SetFloat(float64(n) / 1000)
	case "bool":
		b := false
		if Number[int](1) == 1 {
			b = true
		}
		reflect.ValueOf(&t).Elem().SetBool(b)
	}
	return t
}

func unmarshalStruct(dest any) error {
	rowTypeOf := reflect.TypeOf(dest)
	rowValueOf := reflect.ValueOf(dest)
	if rowTypeOf.Kind() == reflect.Pointer {
		rowTypeOf = rowTypeOf.Elem()
		rowValueOf = rowValueOf.Elem()
	}
	for curField := 0; curField < rowTypeOf.NumField(); curField++ {
		filedName := rowTypeOf.Field(curField).Tag.Get("json")
		if filedName == "" {
			filedName = rowTypeOf.Field(curField).Name
			filedName = toSnake(filedName)
		}
		bind := rowTypeOf.Field(curField).Tag.Get("bind")
		if bind != "" {
			filedName = bind
		}
		for realField, rule := range bindRules {
			curValueOf := rowValueOf.Field(curField)
			curTypeOf := reflect.TypeOf(curValueOf.Interface())
			must := regexp.MustCompile(rule)
			match := must.FindString(filedName)
			if match == "" {
				if curTypeOf.Kind() == reflect.Int {
					value := genField[int]("number")
					rowValueOf.Field(curField).Set(reflect.ValueOf(value))
				}
				if curTypeOf.Kind() == reflect.Int32 {
					value := genField[int32]("number")
					rowValueOf.Field(curField).Set(reflect.ValueOf(value))
				}
				if curTypeOf.Kind() == reflect.Int64 {
					value := genField[int64]("number")
					rowValueOf.Field(curField).Set(reflect.ValueOf(value))
				}
				if curTypeOf.Kind() == reflect.Float32 || curTypeOf.Kind() == reflect.Float32 {
					value := genField[float32]("float")
					rowValueOf.Field(curField).SetFloat(float64(value))
				}
				if curTypeOf.Kind() == reflect.String {
					value := genField[string]("string")
					rowValueOf.Field(curField).SetString(value)
				}
				if curTypeOf.Kind() == reflect.Bool {
					value := genField[bool]("bool")
					rowValueOf.Field(curField).SetBool(value)
				}
				continue
			}
			value := genField[string](realField)
			if curValueOf.Type().Kind() != reflect.String {
				continue
			}
			rowValueOf.Field(curField).SetString(value)
			break
		}
	}
	return nil
}

func toSnake(name string) string {
	var convert []byte
	for i, asc := range name {
		if asc >= 65 && asc <= 90 {
			asc += 32
			if i > 0 {
				convert = append(convert, 95)
			}
		}
		convert = append(convert, uint8(asc))
	}
	return string(convert)
}
