## TDD 过程以及步骤的重要性

编写一个失败的测试，并查看失败信息，我们知道现在有一个为需求编写的 相关 的测试，并且看到它产生了 易于理解的失败描述
编写最少量的代码使其通过，以获得可以运行的程序
然后 重构，基于我们测试的安全性，以确保我们拥有易于使用的精心编写的代码
在我们的例子中，我们通过小巧易懂的步骤从 Hello() 到 Hello("name")，到 Hello("name", "french")。
与「现实世界」的软件相比，这当然是微不足道的，但原则依然通用。TDD 是一门需要通过开发去实践的技能，通过将问题分解成更小的可测试的组件，你编写软件将会更加轻松。
作者：Chris James 译者：Donng 校对：polaris1119，pityonline