function find(arr, target) {
  const lookup = new Map();
  for (let i = 0; i < arr.length; i++) {
    const desired = target - arr[i];
    if (lookup.has(desired)) {
      return [i, lookup.get(desired)];
    } else {
      lookup.set(arr[i], i);
    }
  }
  return [-1, -1];
}
