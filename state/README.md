# 状态模式(State Pattern)

状态模式用于分离状态和行为，允许对象在内部状态发生改变时改变其行为。

将对象的行为封装在不同的状态类中，并通过上下文类来管理状态的切换。 

## 核心思想

将对象的不同状态抽象为独立的类，每个`状态类`都实现了相同的接口。
`上下文类`持有一个当前状态的引用，并将具体的行为委托给当前状态对象处理。
当对象的内部状态发生改变时，上下文类会切换到相应的状态对象。 

## 优点

将复杂的条件判断逻辑转移到不同的状态类中，使代码更加清晰、可维护。
状态模式也符合开闭原则，因为可以通过添加新的状态类来扩展对象的行为，而无需修改现有的代码。

## 缺点

状态模式可能会增加类的数量，因此在使用时需要权衡代码的复杂性和可维护性。

## 适用场景

状态模式适用于对象的行为在不同状态下有所变化的场景，例如订单的不同状态下具有不同的处理逻辑，网络连接的不同状态下执行不同的操作等。 


