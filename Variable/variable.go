package variable

import (
	"fmt"
	_ "math" //加入下畫線 _  導入 可以不使用這個包，這樣可以避免報錯，引用卻不使用
	"strings"
)

// *短變量聲明
func Variable1() {
	myvar1 := 39
	myvar2 := "hello"
	myvar3 := 34.8994959

	fmt.Printf("%d\n", myvar1)
	fmt.Printf("%s\n", myvar2)
	fmt.Printf("%.2f\n", myvar3) //會4捨5入

}

// * 常量變數
func Const() {
	const n = 3.1415926
	const name = "陳宥任"
	const Correct = true
	fmt.Printf("%f\n", n)
	fmt.Printf("常量變數類型:%T\n", n)
	fmt.Printf("%s\n", name)
	fmt.Printf("常量變數類型:%T\n", name)
	fmt.Printf("%v\n", Correct)
	fmt.Printf("常量變數類型:%T\n", Correct)

}

// *格式化輸出
func Datatype() {
	mydata := 64
	fmt.Printf("10進制表示:%d\n", mydata)
	fmt.Printf("mydata數據類型:%T\n", mydata)
	fmt.Printf("8進制表示:%o\n", mydata)
	fmt.Printf("16進制表示:%x\n", mydata)
	fmt.Printf("2進制表示:%b\n", mydata)
	mydata2 := 3.1415926
	fmt.Printf("%.3f\n", mydata2)

}

// *公有/私有變數
/*func Public() {
	var Money = 2000 //首字母是大寫，所以是公有變量
	var xyz = 10     //首字母是小寫，所以是私有的變量
	//xyz :=10 =>短變量無法使用，因為短變量不能使用在全域變數，只能局部

}*/

// *類型轉換
func DataChange() {
	sum := 17
	cont := 5
	var mean float32
	mean = float32(sum) / float32(cont)
	fmt.Printf("%.3f\n", mean)
	//* 精度丟失
	var long1 int64 = 999999999999
	var int1 = int32(long1)
	fmt.Printf("%d\n", int1)

	float1 := 3.026
	int13 := int32(float1) //小數轉成整數，截掉小數部分
	fmt.Printf("%d\n", int13)

}

// *字符串表示
func String() {
	s1 := "hello"
	fmt.Println(s1)
	var s2 = s1 + "world"
	fmt.Println(s2)

	s3 := "\u0048\u0065\u006c\u006f\u0020\u0057\u006f\u0072\u006c\u0064"
	fmt.Println(s3) //unicode編碼
	const s4 = "你好嗎?"
	println(s4)

	//*獲得字符串長度
	fmt.Println(len(s3))
	fmt.Println(s1[2])
	fmt.Println(s3[2])
	fmt.Printf("%c\n", s1[2])
	fmt.Printf("%c\n", s3[2])
}

// *轉義符
func Tag() {
	var s1 = "\"世界\"你好" //加入雙引號""
	fmt.Println(s1)
	var s2 = "\"世界\"\t你好" //加入 制表符
	fmt.Println(s2)
	var s3 = "\"世界\"\\你好" //加入\反協槓
	fmt.Println(s3)
	var s4 = "\"世界\"\n你好" //加入\n換行符
	fmt.Printf("s4=%s", s4)
	var s5 = "\"世界\"\r你好"
	fmt.Printf("s5=%s\n", s5)
	var s6 = "\u0027世界\u0027\r你好"
	fmt.Printf("s5=%s\n", s6)
}

// *字符串應用
func StringUse() {
	string1 := "Hello"
	string2 := "hello"
	result := string1 == string2 //2字串是否相同，相同則true，反之亦然
	println(result)
	text1 := "Go programming"
	substring1 := "Go"
	result = strings.Contains(text1, substring1) //判斷text1字串中 是否 包含substring1
	//result已經聲明過了，所以不用在短變量聲明，只需要賦值
	println(result)

	text2 := "carry"
	fmt.Printf("老字符串:%s\n", text2)
	replacedText := strings.Replace(text2, "r", "t", 1) //用t字串替換text2中的r字串，替換次數n為1
	println(replacedText)
	fmt.Printf("新字符串:%s\n", replacedText)

	text3 := "I love golang"
	splittedString := strings.Split(text3, " ") //按照空格 格式去進行分割
	fmt.Println(splittedString)

	fmt.Println(strings.ToUpper(text3)) //所有字母換成大寫
	fmt.Println(strings.ToLower(text3)) //所有字母換成小寫

}

// *基本運算
func Computation() {
	var x, y = 15, 2

	var result = x + y
	fmt.Printf("x+y=%d\n", result)
	result = x - y
	fmt.Printf("x-y=%d\n", result)
	result2 := x * y
	fmt.Printf("x*y=%d\n", result2)
	result2 = x / y
	fmt.Printf("x/y=%d\n", result2)
	result2 = x % y
	fmt.Println(result2)

	var a, b = 12, 25
	a++
	b--
	// var  c= result2 + (b--)
	//在go語言中，不能這樣做
	fmt.Println(a)
	fmt.Println(b)
}

// *關係運算符
func Equation() {
	var x, y = 15, 25
	fmt.Printf("x==y為: %v\n", x == y)
	fmt.Printf("x!=y為: %v\n", x != y)
	fmt.Printf("x>y為: %v\n", x > y)
	fmt.Printf("x>=y為: %v\n", x >= y)
	fmt.Printf("x<y為: %v\n", x < y)
	fmt.Printf("x<=y為: %v\n", x <= y)
}

// *邏輯運算符
func LogicSign() {
	var x, y, z = 10, 20, 30
	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x < y && x > z))

}

// *位址運算符
func SignPosition() {
	var x, y int = 90, 86
	fmt.Printf("x|y =%b\n", (x | y))
	fmt.Printf("x&y=%b\n", (x & y))
	fmt.Printf("x^y=%b\n", (x ^ y))
	fmt.Printf("x>>2=%b\n", (x >> 2))
	fmt.Printf("x<<2=%b\n", (x << 2))
	fmt.Printf("x>>2=%d\n", (x >> 2))
	fmt.Printf("x<<2=%d\n", (x << 2))
}

// !指標運算
func Ptr() {
	var x int = 100 //聲明變量ｘ
	fmt.Println(x)
	fmt.Printf("x變量的內存地址:%x\n", &x)
	var ptr *int = &x
	fmt.Printf("指針變量ptr值是:%d\n", *ptr) //取出x變量的值

}

// *空指針判斷
func PtrNil() {
	var ptr *int
	fmt.Printf("指針ptr值為%x\n", ptr) //值為 0  =>定義為 空指針
	if ptr == nil {
		fmt.Println("是空指針")
	}
}

// * 二級指針/  指針的指針
func PtrTest() {
	var x int
	var ptr *int   //聲明 int類型指針變量
	var pptr **int //聲明 int類型二級指針變量
	x = 300
	ptr = &x    //獲取x 變量地址
	pptr = &ptr //獲取ptr 變量地址
	fmt.Printf("x:%d\n", x)
	fmt.Printf("*ptr=%d\n", *ptr)
	fmt.Printf("**pptr= %d\n", **pptr)
}

// * 陣列array
func Array() {
	var arr1 = [3]int{1, 2, 3}             //聲明3個元素的int 類型陣列
	var arr2 = [...]float32{1.2, 2.6, 3.6} //#[...]是根據元素長度自動推斷出來的。float32類型陣列
	arr3 := [5]int{4, 5, 6, 7, 8}          //短變量聲明
	arr4 := [...]float32{4, 5, 6.3, 7.6, 5.8}
	fmt.Printf("arr1=%d\n", arr1)
	fmt.Printf("arr2=%f\n", arr2)
	fmt.Printf("arr3=%v\n", arr3)
	fmt.Printf("arr4=%v\n", arr4) //顯示默認數值，所以不會像float32那樣補一堆0

	arr5 := [...]float32{1.2, 2.3, 5.3}
	fmt.Printf("長度: %d~~%f\n", len(arr5), arr5) //計算陣列長度

	arr6 := [5]float32{4.5, 5.1, 6.12}
	fmt.Printf("arr6=%v\n", arr6) //給定陣列長度，但是輸入元素個數卻小於陣列長度，以0去補齊

	arr7 := [3]string{"tony", "Ben"} //字串的默認值，是[ 空字串 ]
	fmt.Printf("arr7=%s\n", arr7)
	fmt.Println(len(arr7))
}

// *訪問陣列
func InvestArray() {
	arr1 := [...]string{"我第1名", "我第2名", "我第3名"}
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1[0]) // 訪問陣列第一個元素

	fmt.Println(arr2[len(arr2)-1]) // 訪問陣列最後一個元素

	var n [10]int
	//賦值
	for i := 0; i < 10; i++ {
		n[i] = i + 100
		fmt.Printf("Element[%v]=%v\n", i, n[i]) //輸出

	}
}

// * 切片 Slice
func Slice() {
	//#  1.Slice在聲明時跟array相同，最大差別為不需要聲明大小，
	//#  2.還有使用make函數作聲明
	strslice1 := []string{"我第1名", "我第2名", "我第3名", "我第4名", "我第5名"}
	intSlice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(strslice1)
	fmt.Println(intSlice1)
	//& make([]T, len, cap) =>T為數據類型，len 為長度，cap為容量可以省略
	var strSlice2 = make([]string, 10, 20)
	fmt.Println(strSlice2)

	var intSlice2 = make([]int, 10)
	fmt.Println(intSlice2)

	fmt.Println(strslice1[0]) // 訪問Slice中的元素
}

// * Slice 切片練習
func SliceTest() {
	numbers := []int{0, 1, 2, 3, 4, 5, 5, 7, 8} //聲明9個元素

	//& Slice切片，不會顯示出上標的值。 Slice[下標:上標]
	fmt.Println(numbers[1:4]) //index 1~4，不包含4，輸出為[1,2,3}
	fmt.Println(numbers[:3])  //下標沒有輸入值，則表示從頭開始，從index [0]開始
	//輸出為[0,1,2]

	fmt.Println(numbers[4:]) //#上標值沒有輸入，則表示輸出到最後元素，而且最後元素也會跟著輸出
	//從index4輸出到最後元素，這個有包含上標值輸出

	a := "hello"
	fmt.Println(a[4:]) //從index4開始到最後元素，但只有o輸出，因為後面沒有元素
	//輸出o

	fmt.Println(a[:3])
	//輸出Hel

	fmt.Println(a[:]) //#前後上標值皆不輸入，直接輸出全部元素

	// fmt.Println(a[:6])      //編譯失敗，因為超出原本的index上標

	fmt.Println(a[:5]) //可以輸出，自訂index上標值最大為5
	b := "applejuice"
	fmt.Println(b[:6])

	//* 追加Slice元素
	//& slice = append(slice, element1 , element2 ,...) => Slice為要被追加的Slice原切片
	var SliceAdd []int
	fmt.Printf("length=%d %d \n", len(SliceAdd), SliceAdd) //長度為0，所以輸出為空
	SliceAdd = append(SliceAdd, 0)                         //追加元素，，輸出[0]
	fmt.Printf("length=%d %d \n", len(SliceAdd), SliceAdd) //
	SliceAdd = append(SliceAdd, 1)                         //追加元素，，輸出[0,1]

	fmt.Printf("length=%d %d \n", len(SliceAdd), SliceAdd)
	SliceAdd = append(SliceAdd, 2, 3, 4, 5)
	fmt.Printf("length=%d %d \n", len(SliceAdd), SliceAdd)
}

// * 映射  map
// &　var 映射變量名[key_data_type]  value_data_type{key:value,key:value,....}
func Map() {
	//#  1.用map關鍵字聲明
	var CountryMap = map[string]string{"TW:": "台灣", "JP": "日本", "USA": "美國"}
	fmt.Printf("length=%v %v\n", len(CountryMap), CountryMap)

	//#  2.通過map() 函數聲明映射
	ClassMap := make(map[int]string)
	fmt.Printf("length=%v %v\n", len(ClassMap), ClassMap) //剛開始make函數，是空的

	ClassMap[101] = "第一名" //&  變數名[key_data] = "value_data"
	ClassMap[102] = "第二名"
	ClassMap[103] = "第三名"
	ClassMap[104] = "第四名"
	ClassMap[105] = "第五名"
	fmt.Printf("length=%v %v\n", len(ClassMap), ClassMap)

	//* 訪問Map元素

	//&  value,result = 映射變量名[key] => 傳回value_data , result是bool值
	name, ok := ClassMap[102]
	if ok {
		fmt.Printf("102學號學生是:%s\n", name)
	} else {
		fmt.Printf("102學號學生不存在!")
	}

	// * 刪除元素

	fmt.Printf("修改前: %v\n長度為:%v\n", ClassMap, len(ClassMap))
	ClassMap[106] = "第10名" //修改value_data
	fmt.Printf("修改後: %v\n長度為:%v\n", ClassMap, len(ClassMap))

	fmt.Printf("刪除前: %v\n長度為:%v\n", ClassMap, len(ClassMap))
	delete(ClassMap, 101) //刪除函數
	fmt.Printf("刪除後: %v\n長度為:%v\n", ClassMap, len(ClassMap))
}

//* For Range使用練習

func ForRange() {
	var CountryMap = map[string]string{"TW:": "台灣", "JP": "日本", "USA": "美國"}
	Odd := [7]int{1, 3, 5, 7, 9, 11, 13}
	for i, value := range Odd {
		fmt.Printf("index:%d  value:%v\n", i, value)
	}

	Strings1 := "Hello world."
	fmt.Println("For range Strings1")
	for i, value := range Strings1 {
		fmt.Printf("index:%d  value:%c\n", i, value)
	}

	fmt.Println("For range Slice CountryMap 輸出:")
	for key, value := range CountryMap {
		fmt.Printf("index:%v value:%v\n", key, value)
	}
}

// * if 練習
func If() {
	fmt.Println("請輸入一個整數:")
	var score int
	fmt.Scan(&score)
	grade := ""
	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if (score >= 60) && (score <= 69) {
		grade = "D"
	} else {
		grade = "E"
	}
	fmt.Printf("分數等級: %s", grade)
}

// * Switch 練習
func Switch() {
	var score int
	grade := ""
	fmt.Println("請輸入<=100的整數:")
	fmt.Scan(&score)
	switch score / 10 {
	case 10:
		grade = "A++"
	case 9:
		grade = "A+"
	case 8:
		grade = "A"
	case 7:
		grade = "B"
	case 6:
		grade = "C"
	default:
		grade = "F"

	}
	fmt.Println("分數等級: " + grade) //用 + 加號， 把字串拼接起來

}

// * Switch 條件多個值 練習
func Switch2() {
	var score int
	grade := ""
	fmt.Println("請輸入<=100的整數:")
	fmt.Scan(&score)
	switch score / 10 {
	case 10, 9: // 10,9先後順序沒有差別， 條件值都是並列同等級的
		grade = "A"

	case 8, 7:
		grade = "B"
	case 6, 5:
		grade = "C"
	default:
		grade = "F"

	}
	fmt.Println("分數等級: " + grade) //用 + 加號， 把字串拼接起來

}

// * Switch  Fallthrough 貫穿case
func SwitchFallthrough() {
	fmt.Println("請輸入一個整數:")
	var season int
	fmt.Scan(&season)
	switch season {
	case 0:
		fmt.Println("春天來了")
	case 1:
		fmt.Println("夏天來了")
	case 2:
		fmt.Println("秋天來了")
		fallthrough //& 貫穿 case，會貫穿而且只會輸出下一個case，其餘csae不會輸出。即輸出case2和case3。
	case 3:
		fmt.Println("秋天天氣最爽啦!")
	default:
		fmt.Println("冬天來了")

	}
}

// * For 迴圈 練習
func For() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d * %d = %d\n", i, i, i*i)
		fmt.Println()
	}
}

// * Break  練習使用
func Break() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, value := range numbers {
		if i == 3 {
			break // 跳出循環
		}
		fmt.Printf("索引: %d     %v\n", i, value)
	}

}

// * 使用標籤的Break 語句
func BreakLabel() {
Outerloop: // 標籤 ，聲明完之後，一定要加入冒號
	for x := 0; x < 5; x++ {
		for y := 5; y > 0; y-- {
			if y == x {
				break Outerloop
			}
			fmt.Printf("(x,y) = (%v,%v)\n", x, y)
		}
	}
}

// * Continue 練習
func Continue() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, value := range numbers {
		if i == 3 { //符合條件的for迴圈條件不輸出
			continue // 不輸出目前的for迴圈條件，繼續執行輸出其他迴圈
		}
		fmt.Printf("Index:%v   value:%v\n", i, value)

	}
}

// * 使用標籤的continue 語句
func ContinueLabel() {
Outerloop:
	for x := 0; x < 5; x++ {
		for y := 5; y > 0; y-- {
			if x == y {
				continue Outerloop
			}
			fmt.Printf("(x,y) = (%v,%v)\n", x, y)

		}
	}
}

// * 使用標籤的GoTo 語句練習
func GotoLabel() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, value := range numbers {
		if i == 3 {
			goto Label //跳到 Label
		}
		fmt.Printf("Index: %v -----value:%v\n", i, value)
	}
Label:
	fmt.Println("執行結束")
}

// * 自定義 Function練習
// func main() {
// 	fmt.Println(Greeting("我第1名"))
// 	fmt.Println(Greeting("我第2名"))
// }

// func Greeting(name string) string { //自定義函數，讓該函數是一個問候函數，參數Name是人名
// 	msg := "嗨!" + name
// 	return msg

// }

// * 自定義函數  return 單一值
func Area(n, w int) int {
	return n * w
}

// * 自定義函數  return 很多值
func ReturnArea(width, height int) (int, int) {

	area := width * height
	perimeter := (width + height) * 2

	return area, perimeter
}

// *自定義函數 不return 值
func CalRectArea(width, height int) (area, perimeter int) { //返回值那一欄位，把變數名字打上去e.g. Area
	//# 返回值欄位已經有了變數名稱，所以只能賦值，不能短變量聲明
	area = width * height //計算面積
	// ar :=width* height
	// area = ar
	perimeter = (width + height) * 2 //計算周長

	//返回數據，第1個值是面積，第2個值是周長
	// return area, perimeter
	return //不返回任何數據類型，因為在返回值欄位已經有名字，所以不能返回任額數據類型
}

// * 自訂義可變參函數
func SumFunction(numbers ...int) int { //& ...表示 形式參數數量不固定，是可變的，同Array  ... 功能
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//* 匿名函數
// func main() {
// 	addFunc := func(a, b int) int { //匿名函數
// 		return a + b
// 	}
// 	subFunc := func(a, b int) int {
// 		return a - b
// 	}
// 	fmt.Printf("%T\n", addFunc)
// 	fmt.Printf("%T\n", subFunc)

// 	var rs func(int, int) int //聲明一個變量，其數據類型為 func(int,int)int
// 	fmt.Println(rs)

// 	var a, b = 10, 5
// 	fmt.Printf("%d\n", addFunc(a, b))
// 	fmt.Printf("%d\n", subFunc(a, b))

// }

// * 函數作為返回值使用
func Calculate(opr string) func(int, int) int {

	var res func(int, int) int //聲明返回值變量

	if opr == "+" {
		res = func(n1, n2 int) int {
			return n1 + n2
		}
	} else {
		res = func(n1, n2 int) int {
			return n1 - n2
		}
	}

	return res
}
