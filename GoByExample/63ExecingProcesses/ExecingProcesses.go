// 在上面的例子中，我们演示了一下如何去触发执行一个外部的进程。
// 我们这样做的原因是我们希望从Go进程里面可以访问外部进程的信息。
// 但有的时候，我们仅仅希望执行一个外部进程来替代当前的Go进程。这个时候，我们需要使用Go提供的exec函数。
package _3ExecingProcesses

import (
    "os"
    "os/exec"
    "syscall"
)

func main() {
    // 本例中，我们使用`ls`来演示。Go需要一个该命令
    // 的完整路径，所以我们使用`exec.LookPath`函数来
    // 找到它
    binary, lookErr := exec.LookPath("adb")
    if lookErr != nil {
        panic(lookErr)
    }
    // `Exec`函数需要一个切片参数，我们给ls命令一些
    // 常见的参数。注意，第一个参数必须是程序名称
    args := []string{"adb"}

    // `Exec`还需要一些环境变量，这里我们提供当前的
    // 系统环境
    env := os.Environ()

    // 这里是`os.Exec`调用。如果一切顺利，我们的原
    // 进程将终止，然后启动一个新的ls进程。如果有
    // 错误发生，我们将获得一个返回值
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}