<h2>Интерфейсы</h2>

<h3>Теория:</h3>

Вопросы

<ul>
    <li>
        Интерфейсы используются для уменьшения связанности кода. Интерфейс должен быть объявлен в месте использования (SOLID).
    </li>
    <li>
        Интерфейс состоит из ссылки на тип данных и ссылки на сами данные.
    </li>
    <li>
        Размер пустого интерфейса = размеру 2 ссылок. В 64-битных системах - 2  * 8 = 16 байт, в 32-битных - 2 * 4 = 8 байт.
    </li>
    <li>
        Переменная типа interface{}/any (пустой интерфейс) будет равна nil, если обе ссылки равны nil, то есть не указан ни тип, ни данные.
    </li>
</ul>



<h3>Задачи:</h3>

<details>
    <summary>
       Лекция 3.2. Nil и не nil интерфейсы. Вариант №1
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        func foo (a interface{}) {
            fmt.Println(a = nil)
        }
        //
        func main() {
            var a int
            var b interface{}
            var c interface{}
            b = 1
            foo(a)
            foo(b)
            foo(c)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func foo (a interface{}) {
                fmt.Println(a = nil)
            }
            //
            func main() {
                var a int
                var b interface{}
                var c interface{}
                b = 1
                foo(a) // false, так как указан тип
                foo(b) // false, так как указаны данные
                foo(c) // true, так как указан пустой тип
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 3.2. Nil и не nil интерфейсы. Объяснить, что выведется и почему, предложить исправленные вариант. Вариант №2
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        type SomeStruct struct{}
        func foo() interface{} {
            var result *SomeStruct 
            return result
        }
        func main() {
            res := foo()
            if res != nil {
                fmt.Println("res != nil, rest =", res)
            }
        }
    </pre></code>
    <details>
        <summary>
            Объяснение
        </summary>
        <pre><code>
            type SomeStruct struct{}
            func foo() interface{} {
                var result *SomeStruct // ссылка на структуру
                return result
            }
            func main() {
                res := foo() // возвращается переменная с типом, который является указателем на пустой интерфейс; != nil, так как указан тип. 
                if res != nil {
                    fmt.Println("res != nil, rest =", res) // res = nil, так как выводится значение
                }
            }
        </pre></code>
    </details>
    <details>
        <summary>
            Исправленный вариант
        </summary>
        <pre><code>
            type SomeInterface interface{}
            func foo() SomeInterface {
                var result SomeInterface
                return result
            }
            func main() {
                res := foo() // возвращается переменная с типом пустой интерфейс; = nil, так как указан тип. 
                if res != nil {
                    fmt.Println("res != nil, rest =", res) // не сработает
                }
            }
        </pre></code>
    </details>
</details>