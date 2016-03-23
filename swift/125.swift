class Solution {
    func filter(c: Character) -> Character {
    if "a" <= c && c <= "z" {
        return c
    }
    
    if "0" <= c && c <= "9" {
        return c
    }
    
    return " "
}
    func isPalindrome(s: String) -> Bool {
        let lower = s.lowercaseString
        var real = ""
        for c in lower.characters {
            let a = filter(c)
            if a != " " {
                real.append(a)
            }
        }
        var reverse = String(real.characters.reverse())
        
        return real == reverse
    }
}

