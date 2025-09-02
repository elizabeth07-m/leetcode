func licenseKeyFormatting(S string, K int) string {
        S = strings.Replace(S, "-", "", -1)
        S = strings.ToUpper(S)
        fmt.Println(S)
        mod := len(S) % K
         fmt.Println(mod)
        if mod == 0 {
            mod += K
        }
        for mod < len(S) {
            S = S[:mod] + "-" + S[mod:]
            mod += K + 1
        }
        return S
}