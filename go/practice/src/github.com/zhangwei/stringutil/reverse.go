// 包stringutil包含用于处理字符串的使用函数。
package stringutil

// 将传入的字符串参数翻转
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
