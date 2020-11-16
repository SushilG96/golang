func defangIPaddr(address string) string {
    res := ""
    for _, i := range address {
        //fmt.Println(i)
        if i == 46{
            res = strings.ReplaceAll(address, ".", "[.]")
        }
    }
    return res
}
