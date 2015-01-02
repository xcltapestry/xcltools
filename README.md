xcltools
========
 用Golang写的一些小东西。 
-------
```C++
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
  
例子：  
代码统计汇总(2014-12-22 22:20:47)
=================================================
分析根目录: E:\GitHub\GitHub\XCL-Charts\XCL-Charts\src

 代码行数     : 文件个数
-------------------------------------------------
 line <= 300  : 92
 line <= 500  : 13
 line <= 1000 : 6
 line <= 5000 : 0
 line > 5000  : 0
-------------------------------------------------
 代码行总计: 18214  注释行总计: 7068
 分析文件数: 111

代码行( >= 600 )文件明细:
   代码行    注释行    文件名
-------------------------------------------------
   987         154    .\org\xclcharts\renderer\AxesChart.java
   620          53    .\org\xclcharts\chart\AreaChart.java
   615          27    .\org\xclcharts\renderer\plot\PlotLegendRender.java
   610          97    .\org\xclcharts\chart\BarChart.java
-------------------------------------------------
             文件数:4

elapsed 1.744222 seconds  
```
