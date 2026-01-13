<h2>Слайсы</h2>

<h3>Теория:</h3>

<ul>
    <li>
        Слайс - динамический массив.
    </li>
    <li>
        Слайс состоит из длины, вместимости, ссылки на массив.
    </li>
    <li>
        Между функциями передается по ссылке.
    </li>
    <li>
        Если копируется слайс, и в него добавляется новый элемент,
        то будет выделена новая память для нового слайса и изменится ссылка при условии, что вместимости не хватает.
    </li>
    <li>
        Если добавлять элементы a = append(a, []int{1, 2, 3, 4, 5}...) (обазятельное нечетное кол-во элементов),
        то длина будет всегда на 1 больше.
    </li>
</ul>


<h3>Задачи:</h3>

<details>
    <summary>
       Лекция 3.5. Вариант №1
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{1, 2, 3}
        fmt.Println("cap(a) = ", cap(a))
        //
        b := append(a, 4)
        c := append(a, 5)
        //
        c[1] = 0
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{1, 2, 3} // l= 3 cap =3
            fmt.Println("cap(a) = ", cap(a)) // c=3
            //
            b := append(a, 4) // l=4, cap=6, new [], b=[1,2,3,4],0,0
            c := append(a, 5) // l=4, cap=6, new [], c=[1,2,3,5],0,0
            //
            c[1] = 0 // c=[1,0,3,5],0,0
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 3.6. Вариант №2
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{}
        a = append(a, []int{1, 2, 3, 4, 5}...)
        fmt.Println("cap(a) = ", cap(a))
        //
        b := append(a, 6)
        c := append(a, 7)
        //
        c[1] = 0
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{} // l=0, c=0
            a= append(a, []int{1, 2, 3, 4, 5}...) // l=5, c=6, a=[1,2,3,4,5],0
            fmt.Println("cap(a) = ", cap(a)) // c=6
            //
            b := append(a, 6) // l=6, c=6, based on a, b=[1,2,3,4,5,6], a=[1,2,3,4,5],6
            c := append(a, 7) // l=6, c=6, based on a, c=[1,2,3,4,5,7], b=[1,2,3,4,5,7], a=[1,2,3,4,5],7
            //
            c[1] = 0 // c=[1,0,3,4,5,7], b=[1,0,3,4,5,7], a=[1,0,3,4,5],7
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 3.7. Вариант №3
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{}
        a = append(a, []int{1, 2, 3, 4, 5}...)
        fmt.Println("cap(a) = ", cap(a))
        //
        b := append(a, 6)
        c := append(b, 7)
        //
        c[1] = 0
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
        //
        d := c[0:12]
        fmt.Println("d =", d)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{} // l=0, c=0
            a = append(a, []int{1, 2, 3, 4, 5}...) // l=5, c=6, a=[1,2,3,4,5],0
            fmt.Println("cap(a) = ", cap(a)) // c=6
            //
            b := append(a, 6) // l=6, c=6, based on a, b=[1,2,3,4,5,6], a=[1,2,3,4,5],6
            c := append(b, 7) // l=7, c=12, new arr, c=[1,2,3,4,5,6,7],0,0,0,0,0
            //
            c[1] = 0 // c=[1,0,3,4,5,6,7],0,0,0,0,0
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
            //
            d := c[0:12]
            fmt.Println("d =", d) // d=[1,0,3,4,5,6,7,0,0,0,0,0]
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.8. Вариант №4
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{}
        for i := range 3 {
            a = append(a, i+1)
        }
        fmt.Println("cap(a) = ", cap(a))
        //
        b := append(a, 4)
        c := append(b, 5)
        //
        c[1] = 0
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
        //
        d := a[0:4]
        fmt.Prinln("d =", d)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{} // l=0, c=0
            for i := range 3 {
                a = append(a, i+1)
            }
            // i=0, l=1, c=1
            // i=1, l=2, c=2
            // i=2, l=3, c=4, a=[1,2,3],0
            fmt.Println("cap(a) = ", cap(a)) // c=4
            //
            b := append(a, 4) // l=4, c=4, based on a, b=[1,2,3,4]
            c := append(b, 5) // l=5, c=8, new array, c=[1,2,3,4,5],0,0,0
            //
            c[1] = 0 // c=[1,0,3,4,5],0,0,0
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
            //
            d := a[0:4]
            fmt.Prinln("d =", d) // a=[1,2,3,4]
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.9. Вариант №5
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{1, 2, 3, 4, 5}
        b := a[2:4]
        c := append(b, 10)
        c[1] = 55
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
        //
        d := b[:3]
        fmt.Println("d =", d)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{1, 2, 3, 4, 5} // l=5, c=5
            b := a[2:4] // l=2, c=3, based on a, b=[3,4],5
            c := append(b, 10) // l=3, c=3, based on b, c=[3,4,10], b=[3,4],10 , a=[1,2,3,4,10]
            c[1] = 55 // c=[3,55,10], b=[3,55],10, a=[1,2,3,55,10]
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
            //
            d := b[:3]
            fmt.Println("d =", d) // d=[3,55,10]
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.10. Вариант №6. Передача по ссылке
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        const MAX = 5
        //
        func main() {
            s: = generate()
            mutation(s)
            fmt.Println("s =", s)
            fmt.Println(s[0:MAX])
        }
        // 
        func generate() []int {
            out := make([]int, 0, MAX)
            for i:= 1; i < MAX; i++ {
                out = append(out, i)
            }
            return out
        }
        //
        func mutation(s []int) {
            s = append(s, -1)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            const MAX = 5
            //
            func main() {
                s: = generate() // l=4, c=5, s=[1,2,3,4],0
                mutation(s) // s=[1,2,3,4],-1
                fmt.Println("s =", s)
                fmt.Println(s[0:MAX]) // s=[1,2,3,4,-1]
            }
            // 
            func generate() []int {
                out := make([]int, 0, MAX) // l=0, c=5, out=[],0,0,0,0,0
                for i:= 1; i < MAX; i++ {
                    out = append(out, i)
                }
                return out
            }
            //
            func mutation(s []int) {
                s = append(s, -1) // l=5, c=5, based on out, s=[1,2,3,4,-1]
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.11. Вариант №7. Задача от GPT
    </summary>
        <p>Исходные данные</p>
        <pre><code>
            s := make([]int, 0, 5)
            s = append(s, 1, 2, 3)
            //
            subSlice := s[1:3]
            //
            subSlice[0] = 99
            subSlice = append(subSlice, 4)
            //
            s = append(s, 5, 6, 7)
            //
            fmt.Println("s = ", s)
            fmt.Println("subSlice = ", subSlice)
        </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            s := make([]int, 0, 5) // l=0, c=5
            s = append(s, 1, 2, 3) // l=3, c=5, s=[1,2,3],0,0
            //
            subSlice := s[1:3] // l=2, c=4, based on s, subSlice=[2,3],0,0
            //
            subSlice[0] = 99 // subSlice=[99,3],0,0, s=[1,99,3],0,0 
            subSlice = append(subSlice, 4) // l=3, c=4, based on s, subSlice=[1,99,3,4],0, s=[1,99,3],4,0 
            //
            s = append(s, 5, 6, 7) // l=6, c=10, ,new array, s=[1,99,3,5,6,7],0,0,0,0
            //
            fmt.Println("s =", s)
            fmt.Println("subSlice =", subSlice)
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.12. Вариант №8. Задача от GPT. Самая сложная
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        func modify(s []int, n int) {
            s = append(s, n)
            s[0] = 999
        }
        //
        func main() {
            s1 := make([]int, 3, 5)
            s2 := s1[:2]
            //
            s1[0] = 1
            s2[1] = 2
            //
            modify(s1, 55)
            modify(s2, 66)
            //
            fmt.Println("s1 =", s1)
            fmt.Println("s2 = ", s2)
            fmt.Println("s1 cap =", cap(s1))
            fmt.Println("s2 cap =", cap(s2))
            //
            s3 :=s2[0:5]
            fmt.Println("s3 =", s3)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
        func modify(s []int, n int) {
            s = append(s, n) // based on s2, s=[999,2,66],55,0, s1=[999,2,66],55,0, s2=[999,2],66,55,0
            s[0] = 999
        }
        //
        func main() {
            s1 := make([]int, 3, 5) // l=3, c=5, s1=[0,0,0],0,0
            s2 := s1[:2] // l=2, c=5, based on s1, s2=[0,0],0,0,0
            //
            s1[0] = 1 // s1=[1,0,0],0,0, s2=[1,0],0,0,0
            s2[1] = 2 // s1=[1,2,0],0,0, s2=[1,2],0,0,0
            //
            modify(s1, 55) // s=[999,2,0,55],0, s1=[999,2,0],55,0, s2[]=[999,2],0,55,0
            modify(s2, 66) // s=[999,2,66],55,0, s1=[999,2,66],55,0, s2=[999,2],66,55,0
            //
            fmt.Println("s1 =", s1)
            fmt.Println("s2 = ", s2)
            fmt.Println("s1 cap =", cap(s1))
            fmt.Println("s2 cap =", cap(s2))
            //
            s3 :=s2[0:5]
            fmt.Println("s3 =", s3) // s2=[999,2,66,55,0]
        }
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 3.13. Вариант №9. Копирование больших данных
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        func getBytes(start, end int) []byte {
            arr := [999999999]byte{} // большой объем данных - около 1 Гб
            slice := arr[start:end]
            return slice // возвращается ссылка
        }
        //
        func main() {
            s := getBytes(10, 20) // работаем с большим слайсом, пока не завершится функция main
            fmt.Println(s)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func getBytes(start, end int) []byte {
                arr := [999999999]byte{}
                //
                slice := make([]byte, end-start)
                copy(slice, arr[start:end])
                //
                return slice 
            }
            //
            func main() {
                s := getBytes(10, 20)
                fmt.Println(s)
            }
        </pre></code>
    </details>
</details>