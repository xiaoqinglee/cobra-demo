go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest


cobra-cli init
cobra-cli add hello_subcommand
cobra-cli add nihao_3rd_level_command -p helloSubcommandCmd


go build
./cobra-demo
./cobra-demo helloSubcommand
./cobra-demo helloSubcommand nihao3rdLevelCommand


Args 是Command的位置参数
Flags 是Command的命名参数


//设置一个当前command和子command的继承Flag
xxxCmd.PersistentFlags().String("name", "", "Say hello to someone")
//设置一个当前command的Flag
xxxCmd.Flags().String("name", "", "Say hello to someone")


//完整flag名 默认值 用法注释
xxxCmd.Flags().String("name", "", "Say hello to someone")
//完整flag名 简写flag名 默认值 用法注释
xxxCmd.Flags().StringP("name", "n", "", "Say hello to someone")


我们可以在Run方法之前或者之后运行一些其它的方法
这些功能按以下顺序运行:

  PersistentPreRun
  PreRun
  Run
  PostRun
  PersistentPostRun

如果子程序没有声明他们自己的该字段的实现, 他们将继承这些功能.
如果子程序声明了他们自己的该字段的实现, 他们自己的该字段的实现会override父命令该字段的实现.


xiaoqing@DESKTOP-I7F5KK1:~/debianEnvGoProj/cobra-demo$ go run main.go helloSubcommand nihao3rdLevelCommand 11 22 -o + -v
nihao3rdLevelCommand called. args are %v [11 22]
11 + 22 = 33
xiaoqing@DESKTOP-I7F5KK1:~/debianEnvGoProj/cobra-demo$ go run main.go helloSubcommand nihao3rdLevelCommand 11 22 -o +
nihao3rdLevelCommand called. args are %v [11 22]
33
xiaoqing@DESKTOP-I7F5KK1:~/debianEnvGoProj/cobra-demo$ go run main.go helloSubcommand nihao3rdLevelCommand 11 22 --operation + --verbose
nihao3rdLevelCommand called. args are %v [11 22]
11 + 22 = 33
