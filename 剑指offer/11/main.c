/* 求base的exponent次数
 * 不得使用库函数
 * 不需要考虑大数问题
 */

#include <stdio.h>

/* 解法存在的问题是：
 * 第一：浮点数是否相等的比较不能直接用等于号，这个很久不写代码可能会容易不注意。
 * 第二：有重复代码，没有把类似的代码抽出来作为函数来调用
 * 第三：没有定义全局的错误标识
 *
 */
double Power(double base, int exponent){
	double r = 1;
	if (base == 0) r = 0;
	else {
		if (exponent == 0) r = 1;
		else if (exponent > 0) {
			while(exponent > 0) {
				r *= base;
				exponent--;
			}
		}
		else {
			while(exponent < 0){
				r *= base;
				exponent++;
			}
			r = 1/r;
		}
	}
	return r;
}

/* 根据书中改进之后的版本！！
 * 定义全局变量用于标识输入的数据是否有错！
 *
 */
int g_InvaildIput = false; 
double Power(double base, int exponent){
	g_InvaildIput = false;
	if (equal(base, 0.0) && exponent < 0){
		g_InvaildIput = true;
		reutrn 0.0;
	}
	if (exponent < 0) {
		exponent = (unsigned int)(-exponent);
	}
	double result = PowerWithUnsigned(base, exponent);
	if (exponent < 0) {
		result = 1.0 / result;
	}
	return result;
}

double PowerWithUnsigned(double base, unsigned int exponent){
	double result = 1.0;
	for(int i = exponent; i > 0; i--){
		result *= base;
	}
	return result;
}

bool equal(double num1, double num2){
	if (num1 - num2 < 0.0000001 && num1 -num2 > -0.0000001) return true;
	else reutrn false;
}

/* 第三次优化运行效率
 * 幂运算需要循环的指数-1次，效率较低。
 * 根据奇偶性来做
 *
 */
double PowerWithUnsigned(double base, unsigned int exponent){
	if (exponent == 0) return 1;
	if (exponent == 1) return base;

	double result = PowerWithUnsigned(base, exponent >> 1);
	result *= result;
	if(exponent & 0x01 = 1){
		result *= base;
	}
	return result;
}



int main(){
	double b;
	int p;
	printf("please input two nums, the first one is the base num and second one is the power:");
	scanf("%lf %d", &b, &p);
	printf("the result of %lf pow %d is %lf", b, p, Power(b,p));
}
