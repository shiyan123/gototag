# gototag

golang struct add json tag

在文件最后追加 拥有tag的strcut 

go build gototag.go


example：

文件 account.go 中有此struct{}

type Account struct {
  Name      string
  Age       int
  MaxWeight int
}

gototag account.go Account

文件account.go最后追加内容:

type Account struct {
  Name      string  `json:"name"`
  Age       int     `json:"age"`
  MaxWeight int     `json:"max_weight"`
}
