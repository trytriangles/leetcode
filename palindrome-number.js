/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
  const xStr = x.toString();
  for (let i = 0; i < xStr.length / 2 + 1; i++) {
    const mirror = xStr[xStr.length - (i + 1)];
    if (xStr[i] !== mirror) {
      return false;
    }
  }
  return true;
};
