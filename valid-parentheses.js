/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
  let stack = [];
  for (let char of s) {
    switch (char) {
      case "(":
        stack.push("(");
        break;
      case ")":
        if (stack.slice(-1)[0] !== "(") return false;
        stack.pop();
        break;
      case "[":
        stack.push("[");
        break;
      case "]":
        if (stack.slice(-1)[0] !== "[") return false;
        stack.pop();
        break;
      case "{":
        stack.push("{");
        break;
      case "}":
        if (stack.slice(-1)[0] !== "{") return false;
        stack.pop();
        break;
    }
  }
  return stack.length === 0;
};
