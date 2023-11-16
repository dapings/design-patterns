# 观察者模式(Observer Pattern)

在多个对象间，定义一个一对多的依赖，当一个对象状态改变时，所有依赖这个对象的对象，都会自动收到通知。

一般情况下，被依赖的对象为被观察者，依赖的对象为观察者。在实际项目中，称呼比较灵活，如 Observable/Observer, Subject/Observer, Publisher/Subscriber, Producer/Consumer, EventEmitter/EventListener, Dispatcher/Listener.

观察者模式用于触发联动，一个对象的改变会触发其它观察者的相关动作，而此对象无需关心连动对象的具体实现。