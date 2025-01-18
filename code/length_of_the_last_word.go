package code

import "strings"

func LengthOfTheLastWord(s string) int {
    s = strings.Trim(s, " ");
    str_arr := strings.Split(s, " ");
    if len(s) == 0 {
        return 0;
    }

    return len(str_arr[len(str_arr) - 1]);
}
