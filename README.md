# 身份证号码验证器

这是一个用于验证和提取身份证号码信息的 Go 语言包。它提供了功能强大的方法来检查身份证号码的有效性，并从中获取出生日期和年龄。

## 功能

- 验证身份证号码的合法性
- 提取身份证号码中的出生日期
- 计算根据出生日期得到的年龄

## 使用方法

### 安装

使用 Go 模块管理，可以通过以下命令安装：

```shell
go get -u github.com/deloz/idcardnumber
```

### 示例

```go
package main

import (
	"fmt"
	
	"github.com/deloz/idcardnumber"
)

func main() {
	id := "320124198701011234" // 这里替换为要验证的身份证号码

	// 验证身份证号码合法性
	err := idcardnumber.Validator(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 创建一个新的身份证号码对象
	card := idcardnumber.New(id)

	// 获取出生日期
	birthday := card.Birthday()
	fmt.Println("出生日期:", birthday)

	// 计算年龄
	age, err := card.Age()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("年龄:", age)
}
```

## 贡献

欢迎贡献代码！如果您有任何建议或发现了 bug，请在 GitHub 上提交 issue 或者提出 pull request。

## 许可证

本项目基于 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。