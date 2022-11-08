#include <stdint.h>

uint32_t reverseBits(uint32_t num) {
    uint32_t result = 0;
    uint32_t shifter = 31;
	while (num != 0) {
        result += (num & 1) << shifter;
        num = num >> 1;
        shifter -= 1;
    }
    return result;
}

int main() {
    uint32_t n = 0b00000010100101000001111010011100;
    uint32_t result = reverseBits(n);
    if (result != 964176192) {
        return 1;
    }
}
