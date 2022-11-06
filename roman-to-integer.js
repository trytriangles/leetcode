const values = new Map([
  ["I", 1],
  ["V", 5],
  ["X", 10],
  ["L", 50],
  ["C", 100],
  ["D", 500],
  ["M", 1000],
]);

/**
 * @param {string} num
 * @return {number}
 */
var romanToInt = function (num) {
  let result = 0;
  let lastValue = 0;
  for (let i = num.length - 1; i >= 0; i--) {
    let value = values.get(num[i]);
    if (value < lastValue) {
      result -= value;
    } else {
      result += value;
    }
    lastValue = value;
  }
  return result;
};
