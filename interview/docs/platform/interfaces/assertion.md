<h2>Интерфейсы</h2>

<h3>Теория:</h3>

<h3>Задачи:</h3>

<details>
    <summary>
        Лекция 3.3. Вариант №1. Изменение типов. Объяснить, что будет выведено
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        type ABC interface {
            A()
            B()
            C()
        }
        type abc struct{}
        //
        func (a abc) A() {}
        func (a abc) B() {}
        func (a abc) C() {}
        //
        type ab struct {}
        //
        func (a ab) A() {}
        func (a ab) B() {}
        //
        func main() {
            // var a = abc{}
            // a1 := a.(ABC)
            // fmt.Println(a1)
            // 
            // var a interface{}
            // a = abc{}
            // a1 := a.(ABC)
            // fmt.Println(a1)
            //
            // var a interface{}
            // a = ab{}
            // a1 := a.(ABC)
            // fmt.Println(a1)
            //
            // var a interface{}
            // a = ab{}
            // a1 := a.(ABC)
            // if !ok {
		    //    fmt.Println("not ok")
	        // }
            // fmt.Println(a1)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            type ABC interface {
                A()
                B()
                C()
            }
            type abc struct{}
            //
            func (a abc) A() {}
            func (a abc) B() {}
            func (a abc) C() {}
            //
            type ab struct {}
            //
            func (a ab) A() {}
            func (a ab) B() {}
            //
            func main() {
                // var a = abc{}
                // a1 := a.(ABC) // так делать нельзя
                // fmt.Println(a1)
                // 
                // var a interface{}
                // a = abc{}
                // a1 := a.(ABC) // сработает
                // fmt.Println(a1)
                //
                // var a interface{}
                // a = ab{}
                // a1 := a.(ABC) // ошибка runtime из-за рефлексии
                // fmt.Println(a1)
                //
                // var a interface{}
                // a = ab{}
                // a1 := a.(ABC) // будет работать
                // if !ok {
                //    fmt.Println("not ok")
                // }
                // fmt.Println(a1)
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.3. Вариант №2. Изменение типов. Объяснить, что будет
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        type ABC interface {
            A()
            B()
            C()
        }
        type AB interface{
            A()
            B()
        }
        type BC interface{
            B()
            C()
        }
        type abc struct{}
        //
        func (a abc) A() {}
        func (a abc) B() {}
        func (a abc) C() {}
        //
        type ab struct {}
        //
        func (a ab) A() {}
        func (a ab) B() {}
        //
        func main() {
            var a interface{}
            a = abc{}
            //
            ab := a.(AB)
            ab.A()
            ab.C()
            //
            bc := ab.(BC)
            bc.C()
            bc.A()
            //
            abc1 := bc.(ABC)
            abc1.A()
            abc1.B()
            abc1.C()
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            type ABC interface {
                A()
                B()
                C()
            }
            type AB interface{
                A()
                B()
            }
            type BC interface{
                B()
                C()
            }
            type abc struct{}
            //
            func (a abc) A() {}
            func (a abc) B() {}
            func (a abc) C() {}
            //
            type ab struct {}
            //
            func (a ab) A() {}
            func (a ab) B() {}
            //
            func main() {
                var a interface{}
                a = abc{} // v=abc, i=ABC
                //
                ab := a.(AB) // v=abc, i=AB
                ab.A() // работает
                ab.C() // ошибка компиляции
                //
                bc := ab.(BC) // v=abc, i=BC
                bc.C() // работает
                bc.A() // ошибка компиляции
                //
                abc1 := bc.(ABC) // v=abc, i=ABC
                abc1.A() // работает
                abc1.B() // работает
                abc1.C() // работает
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.4. Вариант №3. Что будет выведено и как исправить код?
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        func test(x interface{}) {
            switch x.(type) {
            case int:
                fmt.Println("int", x)
            case string:
                fmt.Println("string", x)
            case nil:
                fmt.Println("nil", x)
            case func() int:
                fmt.Prinln("func", x())
            default:
                fmt.Prinln("unknown type")
            }
        }
        func main() {
            var x = func() int {
                return 1
            }
            x = nil
            test(x)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func test(x interface{}) {
                switch x.(type) {
                case int:
                    fmt.Println("int", x)
                case string:
                    fmt.Println("string", x)
                case nil:
                    fmt.Println("nil", x) // x имеет тип - значит != nil
                case func() int:
                    fmt.Println("func", x()) // ошибка компиляции, так тип x = interface{}
                    //
                    f := x.(func() int) // позволит вывести
                    fmt.Println("func", f()) // ошибка компиляции
                default:
                    fmt.Println("unknown type")
                }
            }
            func main() {
                var x = func() int {
                    return 1
                }
                x = nil
                test(x)
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.6. Вариант №4. Как исправить?
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        func readAll(a []io.Reader) {}
        func main() {
            var a []io.ReadWriter{}
            readAll(a)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func convert(a []io.ReaderWriter) []io.Reader {
                var result []io.Reader
                for _, v := range a {
                    result = append(result, v)
                }

                return result
            }
            func readAll(a []io.Reader) {}
            func main() {
                var a []io.ReadWriter{}
                readAll(convert(a))
            }
        </pre></code>
    </details>
</details>