<h2>Интерфейсы</h2>

<h3>Теория:</h3>

<ul>
    <li>
        Интерфейсы используются для уменьшения связанности кода. Интерфейс должен быть объявлен в месте использования (SOLID).
    </li>
    <li>
        <pre><code>
            type iface struct {
                tab *itab
                data unsafe.Pointer
            }
        </code></pre>
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
       Лекция 3.2. Вариант №1. Nil и не nil интерфейсы
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
       Лекция 3.2. Вариант №2. Nil и не nil интерфейсы. Объяснить, что выведется и почему, предложить исправленные вариант
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
                var result *SomeStruct
                return result
            }
            func main() {
                res := foo() // переменная как пустой тип; t = указатель, v = nil
                if res != nil { // true
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
                res := foo() // переменная как пустой тип, t = nil, v = nil
                if res != nil { // false
                    fmt.Println("res != nil, rest =", res)
                }
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 3.2. Вариант №3. Nil и не nil интерфейсы. Объяснить, что выведется и почему, предложить исправленные вариант
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        type SomeError struct{}
        func (s SomeError) Erorr() string {
            return "some error"
        }
        //
        func foo() error {
            var result *SomeError
            return result
        }
        //
        func main() {
            result := foo()
            if result != nil {
                fmt.Println("Error", result)
                return
            }
        }
    </pre></code>
    <details>
        <summary>
            Объяснение
        </summary>
        <pre><code>
            type SomeError struct{}
            func (s SomeError) Erorr() string {
                return "some error"
            }
            //
            func foo() error {
                var result *SomeError
                return result
            }
            //
            func main() {
                result := foo() // t=указатель, v=nil, так как error - это интерфейс
                if result != nil { // true
                    fmt.Println("Error", result) // nil
                    return
                }
            }
        </pre></code>
    </details>
    <details>
        <summary>
            Исправленный вариант. Вариант №1
        </summary>
        <pre><code>
            type SomeError struct{}
            func (s SomeError) Erorr() string {
                return "some error"
            }
            //
            func foo() error {
                var result error
                return result
            }
            //
            func main() {
                result := foo() // t=error, v=Error<nil>
                if result != nil { // true
                    fmt.Println("Error", result) 
                    return
                }
            }
        </pre></code>
    </details>
    <details>
        <summary>
            Исправленный вариант. Вариант №2
        </summary>
        <pre><code>
            var SomeError = errors.New("some error")
            func foo() error {
                var result error
                result = SomeError
                return result
            }
            //
            func main() {
                result := foo() // t=error, v=SomeError
                if result != nil { // true
                    fmt.Println("Error", result) 
                    return
                }
            }
        </pre></code>
    </details>
</details>