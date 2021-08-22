# addOperators
给定一个只含0-9的数字字符串以及一个给定的整数，通过二元操作符（+，-，\*,/）操作字符串中的数字，可得到给定的整数值。试列出所有的二元表达式。注意：单个表达式中，每个数字字符有且只能使用一次。
测试用例：

1、输入"123"以及整数6，可得到["1+2+3","1\*2\*3"]

2、输入"232"以及整数8，可得到["2+2\*3","2\*3+2"]

3、输入"105"以及整数5，可得到["1\*0+5","10-5"]

4、输入"00"以及整数0，可得到["0+0","0-0","0\*0"]

=== RUN   TestAddOperators

  testCase1 [1+2+3 1\*2\*3]
.

1 total assertion

  testCase2 [2+3\*2 2\*3+2]
.

2 total assertions

  testCase3 [1\*0+5 10-5]
.

3 total assertions
  testCase4 [0+0 0-0 0\*0]
.

4 total assertions

--- PASS: TestAddOperators (0.00s)
PASS
ok      _/F_/cyl2021/add/addOperators/src/cmd   0.250s

