package main

//author: XiongChuanLiang
//date:2014-12-22

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/xclpkg/clcolor"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

type fileInfo struct {
	file string
	Code int
	rem  int
}

type InfoSorter []fileInfo

func (info InfoSorter) Len() int {
	return len(info)
}

func (info InfoSorter) Less(i int, j int) bool {
	return info[i].Code > info[j].Code
}

func (info InfoSorter) Swap(i int, j int) {
	info[i], info[j] = info[j], info[i]
}

const GoTimeFormat = "2006-01-02 15:04:05"
const maxGoroutines = 100

func usage() {
	fmt.Println(`
NAME:
  scounter <options> <path> 统计代码行数
OPTIONS:
  -i="": 仅包含指定扩展名的文件(如:.java,.cpp,.h),不输入则包含全部.
  -v=false: 是否显示文件统计明细.
  -l=0: 在统计结果上列出大于等于所指定行数(0为不记录)的文件信息.
EXAMPLE:
  scounter -i .java c:\xclcharts\xclcharts\src
  scounter -i=.cpp,.h,.hpp,.c /xclproject/src
  scounter -i .go -v=false /usr/local/go/src
  scounter -l=680 -i=.cpp,.h,.hpp,.c  /xclproject/common/src
AUTHOR:
  XiongChuanLiang (xcl_168@aliyun.com) `)
}

func main() {
	now := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

	i_flag := flag.String("i", "", "仅包含指定扩展名的文件,逗号分隔(如:.log,.ora),不输入则包含全部文件.")
	v_flag := flag.Bool("v", false, "是否显示文件统计明细.")
	l_flag := flag.Int("l", 0, "在统计结果上列出大于等于所指定行数(0为不记录)的文件信息.")

	flag.Parse()

	if flag.Arg(0) == "" {
		usage()
		os.Exit(1)
	}
	rootPath := filepath.FromSlash(flag.Arg(0))

	fileExt := extSplit(*i_flag)

	infoChan := make(chan fileInfo)

	go findSourceFiles(infoChan, rootPath, fileExt)

	mergeResults(infoChan, *v_flag, *l_flag, rootPath)

	end_time := time.Now()
	var dur_time time.Duration = end_time.Sub(now)
	fmt.Printf("elapsed %f seconds\n", dur_time.Seconds())
}

func mergeResults(infoChan <-chan fileInfo, v_flag bool, l_flag int, rootPath string) {

	var fileNum, code, rem, ln300, ln500, ln1000, ln5000, lnbig int
	var lflagFileInfo InfoSorter

	for info := range infoChan {
		fname := strings.Replace(info.file, rootPath, ".", 1)

		if v_flag {
			fmt.Println("\n file:", fname)
			fmt.Println(" code:", info.Code, " rem:", info.rem)
		}

		if l_flag != 0 && info.Code >= l_flag {
			lflagFileInfo = append(lflagFileInfo, fileInfo{fname, info.Code, info.rem})
		}

		switch {
		case info.Code <= 300:
			ln300++
		case info.Code <= 500:
			ln500++
		case info.Code <= 1000:
			ln1000++
		case info.Code <= 5000:
			ln5000++
		default:
			lnbig++
		}

		code += info.Code
		rem += info.rem
		fileNum++
	}

	title := fmt.Sprintf("\n代码统计汇总(%s)", time.Now().Format(GoTimeFormat))
	fmt.Println(clcolor.Red(title))
	fmt.Println("=================================================")
	fmt.Println("分析根目录:", clcolor.Red(rootPath))
	fmt.Println("\n 代码行数     : 文件个数")
	fmt.Println("-------------------------------------------------")

	fmt.Println(" line <= 300  :", clcolor.Green(toStr(ln300)))
	fmt.Println(" line <= 500  :", clcolor.Green(toStr(ln500)))
	fmt.Println(" line <= 1000 :", clcolor.Red(toStr(ln1000)))
	fmt.Println(" line <= 5000 :", clcolor.Red(toStr(ln5000)))
	fmt.Println(" line > 5000  :", clcolor.Red(toStr(lnbig)))
	fmt.Println("-------------------------------------------------")
	fmt.Println(" 代码行总计:", clcolor.Red(toStr(code)), " 注释行总计:", clcolor.Red(toStr(rem)))
	fmt.Println(" 分析文件数:", clcolor.Red(toStr(fileNum)), "\n")

	if len(lflagFileInfo) > 0 {
		fmt.Println("代码行( >=", clcolor.Blue(toStr(l_flag)), ")文件明细:")
		fmt.Printf("%6s %6s    %s\n", "代码行", "注释行", "文件名")
		fmt.Println("-------------------------------------------------")
		sort.Sort(lflagFileInfo)
		for _, v := range lflagFileInfo {
			fmt.Printf("%6s      %6s    %s\n", clcolor.Blue(toStr(v.Code)), clcolor.Blue(toStr(v.rem)), clcolor.Blue(v.file))
		}
		fmt.Println("-------------------------------------------------")
		fmt.Printf("             文件数:%d\n\n", len(lflagFileInfo))
	}

}

func toStr(v int) string {
	return fmt.Sprintf("%d", v)
}

func extSplit(i_flag string) []string {
	sep := ","
	return strings.Split(i_flag, sep)
}

func findSourceFiles(infoChan chan fileInfo, dirname string, fileext []string) {
	waiter := &sync.WaitGroup{}

	filepath.Walk(dirname, sourceWalkFunc(infoChan, fileext, waiter))
	waiter.Wait()
	close(infoChan)
}

func sourceWalkFunc(infoChan chan fileInfo, fileext []string, waiter *sync.WaitGroup) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err == nil && info.Size() > 0 && !info.IsDir() {

			if isReg(fileext, info.Name()) {
				if runtime.NumGoroutine() > maxGoroutines {
					parseFile(path, infoChan, nil)
				} else {
					waiter.Add(1)
					go parseFile(path, infoChan, func() { waiter.Done() })
				}
			}

		}
		return nil
	}
}

func isReg(fileext []string, filename string) bool {
	var reg bool
	if len(fileext) == 0 {
		reg = true
	} else {
		fext := filepath.Ext(filename)
		for _, ext := range fileext {
			if strings.ToLower(fext) == strings.ToLower(ext) {
				reg = true
			}
		}
	}
	return reg
}

func parseFile(srcfile string, infoChan chan fileInfo, done func()) {
	if done != nil {
		defer done()
	}

	file, err := os.Open(srcfile)
	if err != nil {
		fmt.Println("Failed to open the input file ", file)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	var code, rem int
	var remMulLn bool

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			break
		}

		if isPrefix {
			fmt.Println(srcfile, " : A too long line,seems unexpected.")
			continue
		}

		str := string(line)
		if str == "" {
			continue
		}

		if IsLnRem(str) || IsRem(str) {
			rem++
		} else {
			if strings.Contains(str, "/*") {
				remMulLn = true
				rem++
				continue
			}

			if strings.Contains(str, "*/") {
				remMulLn = false
				rem++
				continue
			}

			if remMulLn {
				rem++
				continue
			}
			code++
		}
	}

	infoChan <- fileInfo{srcfile, code, rem}
}

func IsRem(str string) bool {
	// /* */ /** */
	regular := `^/(\*){1,2}[\W\w]*\*/$`
	return regexp.MustCompile(regular).MatchString(str)
}

func IsLnRem(str string) bool {
	// //
	regular := `^(/){2}[\W\w]*`
	return regexp.MustCompile(regular).MatchString(str)
}
