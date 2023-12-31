# 命令模式(Command Pattern)

命令模式将请求(命令)封装为对象，这样，请求就可以作为参数来传递，并能够支持请求的排队执行、记录日志、撤销等功能。

命令模式的主要优点是它能将请求的调用者和请求的执行者解耦，使得调用者不必知道请求的具体细节，而执行者也不必知道请求的来源。

命令模式的缺点是它会增加系统的复杂度，因为需要引入一个新的类（命令类）。此外，如果命令对象过多，也会影响系统的性能。

命令模式本质是把某个对象的方法调用，封装到对象中，方便传递、存储、调用。

在命令模式中，不同的命令具有不同的目的，对应不同的处理逻辑，并且互相不可替换。

## 应用场景

示例中把主板单中的启动(start)方法和重启(reboot)方法封装为命令对象，再传递到主机(box)对象中，和两个按钮进行绑定：

- 第一个机箱(box1)设置按钮1(button1) 为开机，按钮2(button2)为重启。
- 第二个机箱(box1)设置按钮2(button2) 为开机，按钮1(button1)为重启。

从而得到配置灵活性。

除了配置灵活外，使用命令模式还可以用作：

- 批处理
- 任务队列
- undo, redo
- 在操作系统中，可以使用命令模式来实现文件的复制、删除、移动等操作
- 在游戏中，可以使用命令模式来实现游戏角色的移动、攻击等操作
- 在办公软件中，可以使用命令模式来实现文件的打开、保存、打印等操作

等把具体命令封装到对象中使用的场合。

实际上，Java 中的线程池就用到了命令模式，要执行的逻辑定义在实现了 Runnable 接口的类中，可以实现排队执行、定时执行等。

## 命令模式的实现步骤如下：

1. 创建一个抽象命令类，该类声明了所有命令对象都需要实现的接口。
2. 为每个具体命令创建一个子类，该子类实现了抽象命令类中的接口。
3. 创建一个命令调用者类，该类负责调用命令对象。
4. 在客户端代码中，创建命令对象并将其传递给命令调用者。 