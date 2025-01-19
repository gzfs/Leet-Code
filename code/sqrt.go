package code

func Sqrt(x int) int {
    if x == 0 || x == 1 {
        return 1 * x
    }

    l := 0; 
    r := x/2;

    for l <= r {
        m := l + (r - l) / 2
        if m*m == x {
            return m
        } else {
            if m*m > x {
                r = m - 1
            } else {
                l = m + 1
            }
        }
    }
    return l-1
}
