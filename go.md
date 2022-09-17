##命名
######golang关键字:
    break      default       func     interface   select
    case       defer         go       map         struct
    chan       else          goto     package     switch
    const      fallthrough   if       range       type
    continue   for           import   return      var 
 
######内建常量: 
    true false iota nil
  
######内建类型: 
    int     int8    int16      int32    int64
    uint    uint8   uint16     uint32   uint64 
    float32 float64 complex64  uintptr  complex128
    bool    byte    rune       string   error
  
######内建函数: 
    make len cap new append copy close delete
    complex real imag
    panic recover
##声明
######声明
*    Go语言主要有四种类型的声明语句：var、const、type和func
*    分别对应变量、常量、类型和函数实体对象的声明

##变量

######变量
    *   var 变量名字 类型 = 表达式
    *   其中“类型”或“= 表达式”两个部分可以省略其中的一个
    *   如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。
    *   如果初始化表达式被省略，那么将用零值初始化该变量。
######简短变量声明
    * 在函数内部，简短变量声明语句的形式可用于声明和初始化局部变量。
    * 它以“名字 := 表达式”形式声明变量，变量的类型根据表达式来自动推导。
######new函数
    * 另一个创建变量的方法是调用内建的new函数。
    * 表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T
    * 用new创建变量和普通变量声明语句方式创建变量没有什么区别，除了不需要声明一个临时变量的名字外
######变量的生命周期
    * 变量的生命周期指的是在程序运行期间变量有效存在的时间段。
    * 对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的。
    * 而相比之下，局部变量的生命周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收。
    * 函数的参数变量和返回值变量都是局部变量。它们在函数每次被调用的时候创建。
