#include <stdint.h>
#include <stdio.h>
 
// Using the technique mentioned in K&R C.
int hammingWeight(uint32_t n) {
    int count = 0;
    while (n) {
        n &= (n - 1);
        count++;
    }
    return count;
}
 
int main() {
    if (hammingWeight(11) != 3) {
        return 1;
    }
}