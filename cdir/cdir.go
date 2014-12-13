package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const GoTimeFormat = "2006-01-02 15:04:05"

const (
	MTIME_WIDTH, SIZE_WIDTH = 20, 10
	COL_SEP                 = " "
)

type Options struct {
	T_flag       bool
	S_flag       bool
	D_flag       bool
	F_flag       bool
	A_flag_valid bool
	B_flag_valid bool
	A_flag       time.Time
	B_flag       time.Time
	Include_ext  []string
	Exclude_ext  []string
	TR_flag      bool
	FU_flag      bool
}

var (
	Dir_count, File_count int
)

func usage() {
	fmt.Println(`NAME:
  cdir  显示当前及子目录内容.
OPTIONS:
  -h=false: 显示命令帮助信息
  -a="": 仅显示指定时间(如:2014-10-10_21:14:25)之后的文件或目录.
  -b="": 仅显示指定时间(如:2014-10-10_21:14:25)之前的文件或目录.
  -e="": 指定须排除的指定扩展名文件(如:.bak|.dbf).
  -i="": 仅包含指定扩展名的文件(如:.log|.ora),不输入则包含全部.
  -d=true: 是否显示目录.
  -f=true: 是否显示文件.
  -s=true: 是否显示文件大小.
  -t=true: 是否显示时间.
  -tr=true: 是否以树形方式显示文件或目录.
  -fu=false: 是否以全路径方式显示文件或目录.
EXAMPLE:
  cdir -h 
  ./cdir /usr/local/go 
  cdir -f=false c:\go\doc 
  ./cdir -s=false  /u01/oracle/oradata/xcldb/archivelog -a=2012-11-18_14:27:04
  ./cdir -d=false -fu=true -t=false -e=.out|.go|.jpg|.png /usr/local/go/doc
AUTHOR:
  XiongChuanLiang (xcl_168@aliyun.com) `)
}

func flagParse(options *Options, i_flag, e_flag, a_flag, b_flag string) {
	sep := "|"

	if i_flag != "" {
		options.Include_ext = strings.Split(i_flag, sep)
	} else {
		options.Include_ext = nil
	}

	if e_flag != "" {
		options.Exclude_ext = strings.Split(e_flag, sep)
	} else {
		options.Exclude_ext = nil
	}

	if a_flag != "" {
		A_flag_valid, err := time.Parse(GoTimeFormat, strings.Replace(a_flag, "_", " ", -1))
		if err == nil {
			options.A_flag_valid = true
			options.A_flag = A_flag_valid
		} else {
			fmt.Printf("Error:invalid date:%s\n", a_flag)
			return
		}
	}

	if b_flag != "" {
		B_flag_valid, err := time.Parse(GoTimeFormat, strings.Replace(b_flag, "_", " ", -1))
		if err == nil {
			options.B_flag_valid = true
			options.B_flag = B_flag_valid
		} else {
			fmt.Printf("Error:invalid date:%s\n", b_flag)
			return
		}
	}

}

func main() {
	h_flag := flag.Bool("h", false, "显示命令帮助信息.")
	t_flag := flag.Bool("t", true, "是否显示时间.")
	s_flag := flag.Bool("s", true, "是否显示文件大小.")
	d_flag := flag.Bool("d", true, "是否显示目录.")
	f_flag := flag.Bool("f", true, "是否显示文件.")
	a_flag := flag.String("a", "", "仅显示指定时间(如:2014-10-10_21:14:25)之后的文件或目录..")
	b_flag := flag.String("b", "", "仅显示指定时间(如:2014-10-10_21:14:25)之前的文件或目录..")
	i_flag := flag.String("i", "", "仅包含指定扩展名的文件,逗号分隔(如:.log,.ora),不输入则包含全部文件.")
	e_flag := flag.String("e", "", "指定须排除的指定扩展名文件,逗号分隔(如:.dmp,.dbf).")
	tr_flag := flag.Bool("tr", true, "是否以树形方式显示文件或目录.")
	fu_flag := flag.Bool("fu", false, "是否以全路径方式显示文件或目录.")
	flag.Parse()

	if flag.Arg(0) == "" || *h_flag {
		usage()
		os.Exit(1)
	}
	rootPath := filepath.FromSlash(flag.Arg(0))
	options := &Options{
		A_flag_valid: false,
		B_flag_valid: false,
		T_flag:       *t_flag,
		S_flag:       *s_flag,
		D_flag:       *d_flag,
		F_flag:       *f_flag,
		TR_flag:      *tr_flag,
		FU_flag:      *fu_flag}
	flagParse(options, *i_flag, *e_flag, *a_flag, *b_flag)

	//header
	printHeaderInfo(options)
	//tree
	_, totalSize := dirTree(rootPath, strings.TrimSpace(rootPath), 0, options)
	//total
	printTotalInfo(rootPath, totalSize)
}

func dirTree(basePath, dirName string, level int, options *Options) (bool, int64) {

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error:", err, "\n path:", dirName)
		return false, 0
	}

	var totalSize int64
	level++

	for _, file := range files {
		fullpath := filepath.Join(dirName, file.Name(), GetSeparator())

		if options.B_flag_valid && !file.ModTime().Before(options.B_flag) {
			continue
		}

		if options.A_flag_valid && !file.ModTime().After(options.A_flag) {
			continue
		}

		if file.IsDir() {
			Dir_count++
			printNodeInfo(options, file, level, fullpath)

			_, childTotalSize := dirTree(basePath, fullpath, level, options)
			totalSize += childTotalSize
		} else {
			File_count++
			fExtName := strings.ToLower(filepath.Ext(file.Name()))
			if options.Exclude_ext != nil {
				var bExclude bool = false
				for _, eExt := range options.Exclude_ext {
					if fExtName == eExt {
						bExclude = true
						break
					}
				}
				if bExclude {
					continue
				}
			}

			if options.Include_ext != nil {
				for _, iExt := range options.Include_ext {
					if fExtName == iExt {
						printNodeInfo(options, file, level, fullpath)
						totalSize += file.Size()
						continue
					}
				}
			} else {
				printNodeInfo(options, file, level, fullpath)
				totalSize += file.Size()
			}

		}
	}
	return true, totalSize
}

func printNodeInfo(options *Options, file os.FileInfo, level int, fullpath string) {

	var (
		nodeName, mtime, fsize string
		nodeDir                bool
	)

	if (file.Mode() & os.ModeDir) > 0 {
		if !options.D_flag {
			return
		}
		nodeDir = true
	} else {
		if !options.F_flag {
			return
		}
	}

	if options.FU_flag {
		nodeName = fullpath
	} else {
		nodeName = file.Name()
	}

	if !nodeDir && runtime.GOOS != "windows" {
		mode := GetFileModeString(file.Mode())
		if mode != "" {
			nodeName = "[" + mode + "] " + nodeName
		}
	}

	if nodeDir {
		nodeName = " + " + nodeName
		fsize = strings.Repeat(" ", SIZE_WIDTH)
	} else {
		nodeName = "   " + nodeName
		fsize = GetSizeString(file.Size(), SIZE_WIDTH)
	}

	if options.TR_flag {
		nodeName = strings.Repeat(" ", 2*level) + nodeName
	}

	mtime = file.ModTime().Format(GoTimeFormat)

	switch {
	case options.T_flag && options.S_flag:
		fmt.Println(mtime, COL_SEP, fsize, nodeName)
	case options.T_flag:
		fmt.Println(mtime, nodeName)
	case options.S_flag:
		fmt.Println(fsize, nodeName)
	default:
		fmt.Println(nodeName)
	}
}

func printHeaderInfo(options *Options) {

	mt := strings.Repeat(" ", MTIME_WIDTH-len("MTime"))
	sz := strings.Repeat(" ", SIZE_WIDTH-len("Size"))

	mtln := strings.Repeat("-", MTIME_WIDTH)
	szln := strings.Repeat("-", SIZE_WIDTH)

	switch {
	case options.T_flag && options.S_flag:
		fmt.Println("MTime", mt, COL_SEP, "Size", sz)
		fmt.Println(mtln, COL_SEP, szln)
	case options.T_flag:
		fmt.Println("MTime", mt)
		fmt.Println(mtln)
	case options.S_flag:
		fmt.Println("Size", sz)
		fmt.Println(szln)
	default:
		fmt.Println(szln)
	}
}

func printTotalInfo(rootPath string, totalSize int64) {
	fmt.Println(strings.Repeat("-", MTIME_WIDTH+SIZE_WIDTH+4))
	fmt.Println(rootPath)
	fmt.Println(" ", Dir_count, " directories,", File_count, " files,",
		GetSizeString(totalSize, SIZE_WIDTH), "(", totalSize, " B)")
}

func GetSizeString(size int64, maxWidth int) string {

	var kb int64 = 1024
	var mb int64 = kb * 1024
	var gb int64 = mb * 1024
	var tb int64 = gb * 1024

	var sizeStr string

	if size < kb {
		sizeStr = fmt.Sprintf("%d B", size)
	} else if size < mb {
		sizeStr = fmt.Sprintf("%d KB", size/kb)
	} else if size < gb {
		sizeStr = fmt.Sprintf("%d MB", size/mb)
	} else if size < tb {
		sizeStr = fmt.Sprintf("%d GB", size/gb)
	} else {
		sizeStr = fmt.Sprintf("%d TB", size/tb)
	}
	sizeLen := len(sizeStr)

	if -1 != maxWidth || maxWidth > sizeLen {
		sizeStr = strings.Repeat(" ", maxWidth-sizeLen) + sizeStr
	}
	return sizeStr
}

func GetSeparator() string {
	return string(filepath.Separator)
}

func GetFileModeString(fMode os.FileMode) string {

	if (fMode & os.ModeSymlink) > 0 {
		return "L" // L: symbolic link
	}

	if (fMode & os.ModeNamedPipe) > 0 {
		return "P" // p: named pipe (FIFO)
	}

	if (fMode & os.ModeSocket) > 0 {
		return "S" // S: Unix domain socket
	}

	if (fMode & os.ModeDevice) > 0 {
		return "D" // D: device file
	}

	return ""
}
