package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    var s string
    fmt.Scan(&s)
    
    count := 0
    
    for i := 0; i < n; i++ {
        if isVowel(s[i]) {
            count += n - i // add all substrings ending with a vowel
            for j := i+1; j < n; j++ {
                if isVowel(s[j]) {
                    count -= n - j // remove substrings ending with a vowel after this vowel
                    break
                }
            }
        }
    }
    
    fmt.Println(count)
}

func isVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
