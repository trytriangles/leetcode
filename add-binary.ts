function addBinary(a: string, b: string): string {
  const bigA = BigInt("0b" + a);
  const bigB = BigInt("0b" + b);
  return (bigA + bigB).toString(2);
}
