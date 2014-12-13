//author:xiong chanliang
//date: 2014-12-13
#include <iostream>
#include <cstdlib>
#include <tuple>
#include <vector>
#include <map>
#include <algorithm>
#include <numeric> //iota

using namespace std;


///////////////////////////////////////////////////
// 初始化
///////////////////////////////////////////////////
void init_demo(){
	//初始化
	int mArr[] = { 1, 2, 3 };
	auto mList = vector < int > {1, 2, 3, 4};
	auto mMap = map < int, string > {{1, "aa"}, { 2, "bb" }};

	cout << "vector: ";
	for (const int& x : mList)
		cout << x << " ";
	cout << endl;

	cout << "map: ";
	for (const auto& mp : mMap)
		cout << mp.first << " " << (mp.second).c_str();
	cout << endl;
}


///////////////////////////////////////////////////
// raw字符串字面值
///////////////////////////////////////////////////
void raw_demo(){

	string path = R"(c:\a\b\c\GG再也不用烦转义符了)";
	cout << "单行:" << path.c_str() << endl;

	string muln = R"(C++/Golang
	 aa'aaa\Cplusplus/gogogo
	 author"xiongchuanliang
	)";
	cout << "多行:" << muln.c_str() << endl;
}


///////////////////////////////////////////////////
// 值顺序递增
///////////////////////////////////////////////////
void iota_demo(){
	int d[5] = { 0 };
	std::iota(d, d + 5, 10);

	cout << "iota_demo(): old : d[5] = { 0 } "<< endl;
	cout << "iota_demo(): iota: d[5] = { ";
	for each (int var in d)
	{
		cout <<  var <<" ";
	}
	cout <<"} "<< endl;

	char e[5] = { 'a' };
	char f[5] = { 0 };
	copy_n(e, 5, f);

	cout << "iota_demo(): old : e[5] = { 'a' } " << endl;
	cout << "iota_demo(): iota: e[5]  " << endl;
	std::iota(e, e + 5, 'e');
	for (size_t i = 0; i < 5; i++)
	{
		cout << " iota = " << e[i] << endl;
	}
	

}

///////////////////////////////////////////////////
// lambda
///////////////////////////////////////////////////
void lambda_demo(){
	int arr[] = { 1, 2, 3 };
	for_each(begin(arr), end(arr), [](int n){cout << n << endl; });

	cout << "---lambda_demo---" << endl;

	auto func = [](int n){cout << n << endl; };
	for_each(begin(arr), end(arr), func);
}

///////////////////////////////////////////////////
// 多值赋值
///////////////////////////////////////////////////
tuple<int, string> tuple_demo(){
	tuple<int, string> ret;
	ret = make_tuple(168, "函数返回的字符串");
	cout << "tuple_demo(): " << get<0>(ret) << " " << (get<1>(ret)).c_str() << endl;

	auto triple = make_tuple(5, 6, 7);
	cout << "tuple_demo(): " << get<0>(triple) << " " << get<1>(triple) << " " << get<2>(triple) << endl;

	int ti;
	string ts;
	tie(ti, ts) = make_tuple(10, "xcl--将数字和字符赋值给两个变量");
	cout << "tuple_demo(): " << ti << " " << ts.c_str() << endl;

	return ret;
}


///////////////////////////////////////////////////
// 变长参数
///////////////////////////////////////////////////
template<typename T> void fmtPrintln(T value){
	cout << value << endl;
}

template<typename T, typename... Args>
void fmtPrintln(T head, Args... args)
{
	cout << head << " ";
	fmtPrintln(args...);
}


void fmtPrintf(const char *s)
{
	while (*s) {
		if (*s == '%') {
			if (*(s + 1) == '%') {
				++s;
			}
			else {
				throw std::runtime_error("invalid format string: missing arguments");
			}
		}
		std::cout << *s++;
	}
}

template<typename T, typename... Args>
void fmtPrintf(const char *s, T value, Args... args)
{	
	while (*s) {
		if (*s == '%') {
			if (*(s + 1) == '%') {
				++s;
			}
			else {
				std::cout << value;
				// call even when *s == 0 to detect extra arguments
				fmtPrintf(s + 1, args...);
				return;
			}
		}
		std::cout << *s++;
	}
	throw std::logic_error("extra arguments provided to printf");
}


template<typename... Args> int countArgs(Args... args){
	return sizeof...(args);
}
///////////////////////////////////////////////////



///////////////////////////////////////////////////
// map
///////////////////////////////////////////////////
typedef map <string, int > map_str_int;
tuple<string, int, bool> mapfind_demo(map_str_int myMap, string key){
	//typedef map <string, int > map_str_int;
	//auto myMap = map_str_int{ { "aa", 1 }, { "bb", 2 }, { "cc", 3 } };

	map_str_int::iterator pos;	
	pos = myMap.find(key);
	if (pos == myMap.end()){
		return make_tuple("", 0, false);
	}
	else{
		return make_tuple(pos->first, pos->second, true);
	}
}

/*
auto found = [=](const map_str_int::value_type &mp){
if (mp.first == key)
{
cout << "found map=" << mp.second << endl;
return true;
}
return false;
};
//pos = find_if(begin(myMap), end(myMap), found);
*/


///////////////////////////////////////////////////
// 回调函数
///////////////////////////////////////////////////
void printFunc(string arg){
	cout << "callFunc() -> printFunc():" << arg.c_str() << endl;
}

typedef void(*callf)(string);
void callFunc(callf pFunc, string arg){ //void(*pFunc)(string)
	pFunc(arg);
}
//callFunc(printFunc, "回调就是你调我，我调它，大家一起玩。");
///////////////////////////////////////////////////
// 泛型
///////////////////////////////////////////////////
template <typename T> int compare(const T v1, const T v2){
	if (v1 < v2){
		cout << "compare(): v1 < v2" << endl;
		return -1;
	}
	else if (v1 == v2){
		cout << "compare(): v1 == v2" << endl;
		return 0;
	}
	else{
		cout << "compare(): v1 > v2" << endl;
		return 1;
	}
}

///////////////////////////////////////////////////
// slice
///////////////////////////////////////////////////
void slice_demo(){

	int a[5] = { 1, 2, 3, 4, 5 };
	int b[3] = { 0 };
	int c[2] = { 0 };


	cout << "a[5] = { 1, 2, 3, 4, 5 }" << endl;
	cout << "array[:end_pos]: b = array[:3]" << endl;
	// array[:end_pos]
	copy_n(a, 3, b);
	for each (int var in b)
	{
		cout << " " << var;
	}
	cout << endl;


	cout << "a[5] = { 1, 2, 3, 4, 5 }" << endl;
	cout << "array[begin_pos:end_pos]: c = array[1,2] " << endl;
	// array[begin_pos:end_pos]
	int begin_pos = 1;
	int subLen = sizeof(c) / sizeof(c[0]);
	int end_pos = begin_pos + subLen;
	copy_if(a + begin_pos, a + end_pos, c, [](int v){return true; });
	
	for each (int var in c)
	{
		cout << " " << var ;
	}
	cout << endl;
}



int main(void){

	cout << "Golang部份特性的C++实现对比." << endl;

	cout << "\nraw字符串字面值 " << endl;
		raw_demo();
	cout << "raw字符串字面值 end." << endl;

	cout << "\n变量初始化 " << endl;
		init_demo();
	cout << "变量初始化 end." << endl;
	
	cout << "\nlambda " << endl;
		lambda_demo();
	cout << "lambda end." << endl;	
	
	cout << "\n值顺序递增 " << endl;
		iota_demo();
	cout << "值顺序递增 end." << endl;

	cout << "\n多值赋值及函数返回多值 " << endl;
		int ti;
		string ts;
		tie(ti, ts) = tuple_demo();
		cout << "main() <- tuple_demo(): " << ti << " " << ts.c_str() << endl;
	cout << "多值赋值及函数返回多值 end." << endl;

	cout << "\nmap查找 " << endl;
		auto myMap = map_str_int{ { "aa", 1 }, { "bb", 2 }, { "cc", 3 } };

		string mpKey;
		int mpValue;
		bool  mpFound = false;
		tie(mpKey, mpValue, mpFound) = mapfind_demo(myMap, "bb");

		if (mpFound){
			cout << "mapfind_demo: found" << endl;
		}
		else{
			cout << "mapfind_demo: not found" << endl;
		}
	cout << "函数返回多值/map查找 end." << endl;
	
	cout << "\n变长参数 " << endl;
		fmtPrintln("fmtPrintln(): ", 1, 2.0, "C++11", "Golang");

		fmtPrintf(":: %s,%d,%f,%s,%s \n", "fmtPrintf(): ", 1, 2.0, "C++11", "Golang");
	cout << "变长参数 end." << endl;

	cout << "\n回调函数 " << endl;
		callFunc(printFunc, "回调就是你调我，我调它，大家一起玩。");
	cout << "回调函数 end." << endl;

	cout << "\n泛型 " << endl;
		int i1 = 5, i2 = 7;
		double d1 = 52.5, d2 = 10.7;
		compare(i1, i2);
		compare(d1, d2);
	cout << "泛型 end." << endl;

	cout << "\n数组和切片(sclie) " << endl;
		slice_demo();
	cout << "数组和切片(sclie) end." << endl;

	system("pause");
}
