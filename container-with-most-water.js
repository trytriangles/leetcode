/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  let i = 0;
  let j = height.length - 1;
  let previousI = 0;
  let previousJ = 0;
  let result = 0;
  while (i < j) {
    const width = j - i;
    const length = Math.min(height[j], height[i]);
    const area = length * width;
    if (area > result) result = area;
    if (height[i] < height[j]) {
      while (height[i] <= previousI) i++;
      previousI = height[i];
    } else {
      while (height[j] <= previousJ) j--;
      previousJ = height[j];
    }
  }
  return result;
};
