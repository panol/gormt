package cnf

// EImportsHead imports head options. import包含选项
var EImportsHead = map[string]string{
	"stirng":         `"string"`,
	"time.Time":      `"time"`,
	"gorm.Model":     `"gorm.io/gorm"`,
	"fmt":            `"fmt"`,
	"datatypes.JSON": `"gorm.io/datatypes"`,
	"datatypes.Date": `"gorm.io/datatypes"`,
}

// TypeMysqlDicMp Accurate matching type.精确匹配类型
var TypeMysqlDicMp = map[string]string{
	"smallint":            "int32",
	"smallint unsigned":   "uint32",
	"int":                 "int32",
	"int unsigned":        "uint32",
	"bigint":              "int64",
	"bigint unsigned":     "uint64",
	"varchar":             "string",
	"char":                "string",
	"date":                "datatypes.Date",
	"datetime":            "time.Time",
	"bit(1)":              "[]uint8",
	"tinyint":             "int32",
	"tinyint unsigned":    "uint32",
	"tinyint(1)":          "int32",  // tinyint(1) 默认设置成bool
	"tinyint(1) unsigned": "uint32", // tinyint(1) 默认设置成bool
	"json":                "datatypes.JSON",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"double unsigned":     "float64",
	"mediumtext":          "string",
	"longtext":            "string",
	"float":               "float32",
	"float unsigned":      "float32",
	"tinytext":            "string",
	"enum":                "string",
	"time":                "time.Time",
	"tinyblob":            "[]byte",
	"blob":                "[]byte",
	"mediumblob":          "[]byte",
	"longblob":            "[]byte",
	"integer":             "int64",
}

// TypeMysqlMatchList Fuzzy Matching Types.模糊匹配类型
var TypeMysqlMatchList = []struct {
	Key   string
	Value string
}{
	{`^(tinyint)[(]\d+[)] unsigned`, "uint32"},
	{`^(smallint)[(]\d+[)] unsigned`, "uint32"},
	{`^(int)[(]\d+[)] unsigned`, "uint32"},
	{`^(bigint)[(]\d+[)] unsigned`, "uint64"},
	{`^(float)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(double)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(tinyint)[(]\d+[)]`, "int32"},
	{`^(smallint)[(]\d+[)]`, "int32"},
	{`^(int)[(]\d+[)]`, "int32"},
	{`^(bigint)[(]\d+[)]`, "int64"},
	{`^(char)[(]\d+[)]`, "string"},
	{`^(enum)[(](.)+[)]`, "string"},
	{`^(varchar)[(]\d+[)]`, "string"},
	{`^(varbinary)[(]\d+[)]`, "[]byte"},
	{`^(blob)[(]\d+[)]`, "[]byte"},
	{`^(binary)[(]\d+[)]`, "[]byte"},
	{`^(decimal)[(]\d+,\d+[)]`, "float64"},
	{`^(mediumint)[(]\d+[)]`, "string"},
	{`^(double)[(]\d+,\d+[)]`, "float64"},
	{`^(float)[(]\d+,\d+[)]`, "float64"},
	{`^(datetime)[(]\d+[)]`, "time.Time"},
	{`^(bit)[(]\d+[)]`, "[]uint8"},
	{`^(text)[(]\d+[)]`, "string"},
	{`^(integer)[(]\d+[)]`, "int32"},
}
