/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function (nums, val) {
  const indexesToRemove = [];
  nums.forEach((num, index) => {
    if (num === val) indexesToRemove.push(index);
  });
  indexesToRemove.forEach((indexToRemove, offset) => {
    nums.splice(indexToRemove - offset, 1);
  });
  return nums.length;
};
