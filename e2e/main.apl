import lib
import lib2

func foo() int {
    int i = 5;
    return i; 
}

func bar(int x) {
    if (x == 1) {
        print(x);
    } else {
        print("not 1");
    }
}

func main() {
    bar(foo() + lib.foo() + lib2.foo())
}
